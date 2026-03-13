package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const schemaTable = "schema_migrations"

// Direction indicates whether to apply or revert migrations.
type Direction string

const (
	Up   Direction = "up"
	Down Direction = "down"
)

// Migration represents a single versioned migration file pair.
type Migration struct {
	Version  int
	Name     string
	UpFile   string
	DownFile string
}

// Runner executes SQL migrations from a directory against a database.
type Runner struct {
	db           *sql.DB
	migrationsFS fs.FS  // optional embedded FS; nil means use disk
	dir          string // disk path when migrationsFS is nil
}

// NewRunner creates a Runner that reads migration files from dir on disk.
// The directory is discovered automatically from the environment variable
// MIGRATION_DIR, falling back to the project-relative default.
func NewRunner(db *sql.DB, dir string) *Runner {
	if envDir := os.Getenv("MIGRATION_DIR"); envDir != "" {
		dir = envDir
	}
	return &Runner{db: db, dir: dir}
}

// NewRunnerFS creates a Runner backed by an embedded fs.FS (for embedded migrations).
func NewRunnerFS(db *sql.DB, migFS fs.FS) *Runner {
	return &Runner{db: db, migrationsFS: migFS}
}

// =====================================================================
// Public API
// =====================================================================

// Up applies all pending migrations.
func (r *Runner) Up() error {
	if err := r.ensureSchemaTable(); err != nil {
		return err
	}
	applied, err := r.appliedVersions()
	if err != nil {
		return err
	}
	all, err := r.loadMigrations()
	if err != nil {
		return err
	}

	pending := filterPending(all, applied)
	if len(pending) == 0 {
		log.Println("No pending migrations.")
		return nil
	}
	for _, m := range pending {
		if err := r.apply(m); err != nil {
			return fmt.Errorf("migration %d_%s failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}

// Down rolls back the last N migrations (default 1).
func (r *Runner) Down(steps int) error {
	if steps <= 0 {
		steps = 1
	}
	if err := r.ensureSchemaTable(); err != nil {
		return err
	}
	applied, err := r.appliedVersions()
	if err != nil {
		return err
	}
	all, err := r.loadMigrations()
	if err != nil {
		return err
	}

	toRevert := appliedInReverseOrder(all, applied, steps)
	if len(toRevert) == 0 {
		log.Println("Nothing to roll back.")
		return nil
	}
	for _, m := range toRevert {
		if err := r.revert(m); err != nil {
			return fmt.Errorf("rollback %d_%s failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}

// Status prints which migrations have been applied and which are pending.
func (r *Runner) Status() error {
	if err := r.ensureSchemaTable(); err != nil {
		return err
	}
	applied, err := r.appliedVersions()
	if err != nil {
		return err
	}
	all, err := r.loadMigrations()
	if err != nil {
		return err
	}

	fmt.Printf("%-8s %-40s %-10s\n", "VERSION", "NAME", "STATUS")
	fmt.Println(strings.Repeat("-", 62))
	for _, m := range all {
		status := "pending"
		if _, ok := applied[m.Version]; ok {
			status = "applied"
		}
		fmt.Printf("%-8d %-40s %-10s\n", m.Version, m.Name, status)
	}
	return nil
}

// Version returns the currently applied migration version (highest applied).
func (r *Runner) Version() (int, error) {
	if err := r.ensureSchemaTable(); err != nil {
		return 0, err
	}
	applied, err := r.appliedVersions()
	if err != nil {
		return 0, err
	}
	max := 0
	for v := range applied {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Force marks a specific version as applied without running SQL.
// Use only for recovering from a dirty state in production.
func (r *Runner) Force(version int) error {
	if err := r.ensureSchemaTable(); err != nil {
		return err
	}
	_, err := r.db.Exec(
		fmt.Sprintf(`INSERT INTO %s (version, applied_at) VALUES ($1, $2)
		             ON CONFLICT (version) DO UPDATE SET applied_at = EXCLUDED.applied_at`, schemaTable),
		version, time.Now().UTC(),
	)
	if err != nil {
		return fmt.Errorf("force version %d: %w", version, err)
	}
	log.Printf("Forced migration version to %d", version)
	return nil
}

// Drop reverts ALL migrations in reverse order (wipes the schema).
func (r *Runner) Drop() error {
	if err := r.ensureSchemaTable(); err != nil {
		return err
	}
	applied, err := r.appliedVersions()
	if err != nil {
		return err
	}
	all, err := r.loadMigrations()
	if err != nil {
		return err
	}
	toRevert := appliedInReverseOrder(all, applied, len(all))
	for _, m := range toRevert {
		if err := r.revert(m); err != nil {
			return fmt.Errorf("drop/rollback %d_%s failed: %w", m.Version, m.Name, err)
		}
	}
	return nil
}

// =====================================================================
// Internal helpers
// =====================================================================

func (r *Runner) ensureSchemaTable() error {
	_, err := r.db.Exec(fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version    INTEGER      PRIMARY KEY,
			applied_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
		)`, schemaTable))
	return err
}

func (r *Runner) appliedVersions() (map[int]time.Time, error) {
	rows, err := r.db.Query(fmt.Sprintf(`SELECT version, applied_at FROM %s ORDER BY version`, schemaTable))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[int]time.Time)
	for rows.Next() {
		var v int
		var t time.Time
		if err := rows.Scan(&v, &t); err != nil {
			return nil, err
		}
		result[v] = t
	}
	return result, rows.Err()
}

// loadMigrations discovers all *.up.sql / *.down.sql pairs and returns them
// sorted ascending by version number. New files dropped in the directory are
// automatically picked up without any code change.
func (r *Runner) loadMigrations() ([]Migration, error) {
	entries, err := r.readDir(".")
	if err != nil {
		return nil, fmt.Errorf("reading migration dir: %w", err)
	}

	// map version -> Migration
	byVersion := make(map[int]*Migration)
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".sql") {
			continue
		}
		ver, migName, direction, ok := parseMigrationFilename(name)
		if !ok {
			continue
		}
		m, exists := byVersion[ver]
		if !exists {
			m = &Migration{Version: ver, Name: migName}
			byVersion[ver] = m
		}
		if direction == string(Up) {
			m.UpFile = name
		} else {
			m.DownFile = name
		}
	}

	migrations := make([]Migration, 0, len(byVersion))
	for _, m := range byVersion {
		if m.UpFile == "" {
			return nil, fmt.Errorf("migration %d (%s) has no .up.sql file", m.Version, m.Name)
		}
		migrations = append(migrations, *m)
	}
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})
	return migrations, nil
}

func (r *Runner) apply(m Migration) error {
	sql, err := r.readFile(m.UpFile)
	if err != nil {
		return err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(sql); err != nil {
		_ = tx.Rollback()
		return err
	}
	if _, err := tx.Exec(
		fmt.Sprintf(`INSERT INTO %s (version, applied_at) VALUES ($1, $2)`, schemaTable),
		m.Version, time.Now().UTC(),
	); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	log.Printf("[UP]   %06d_%s", m.Version, m.Name)
	return nil
}

func (r *Runner) revert(m Migration) error {
	if m.DownFile == "" {
		return fmt.Errorf("migration %d (%s) has no .down.sql file; cannot roll back", m.Version, m.Name)
	}
	sql, err := r.readFile(m.DownFile)
	if err != nil {
		return err
	}
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(sql); err != nil {
		_ = tx.Rollback()
		return err
	}
	if _, err := tx.Exec(
		fmt.Sprintf(`DELETE FROM %s WHERE version = $1`, schemaTable),
		m.Version,
	); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	log.Printf("[DOWN] %06d_%s", m.Version, m.Name)
	return nil
}

// readDir lists directory entries from the embedded FS or from disk.
func (r *Runner) readDir(path string) ([]fs.DirEntry, error) {
	if r.migrationsFS != nil {
		return fs.ReadDir(r.migrationsFS, path)
	}
	return os.ReadDir(r.dir)
}

// readFile reads a migration SQL file from the embedded FS or from disk.
func (r *Runner) readFile(name string) (string, error) {
	var data []byte
	var err error
	if r.migrationsFS != nil {
		data, err = fs.ReadFile(r.migrationsFS, name)
	} else {
		data, err = os.ReadFile(filepath.Join(r.dir, name))
	}
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// =====================================================================
// Pure helpers
// =====================================================================

// parseMigrationFilename parses "000001_create_users_table.up.sql"
// into (1, "create_users_table", "up", true).
func parseMigrationFilename(name string) (version int, migName, direction string, ok bool) {
	// strip .sql
	withoutSQL := strings.TrimSuffix(name, ".sql")
	// must end in .up or .down
	var suffix string
	if strings.HasSuffix(withoutSQL, ".up") {
		suffix = "up"
		withoutSQL = strings.TrimSuffix(withoutSQL, ".up")
	} else if strings.HasSuffix(withoutSQL, ".down") {
		suffix = "down"
		withoutSQL = strings.TrimSuffix(withoutSQL, ".down")
	} else {
		return 0, "", "", false
	}
	// remainder: "000001_create_users_table"
	parts := strings.SplitN(withoutSQL, "_", 2)
	if len(parts) != 2 {
		return 0, "", "", false
	}
	v, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", "", false
	}
	return v, parts[1], suffix, true
}

func filterPending(all []Migration, applied map[int]time.Time) []Migration {
	var pending []Migration
	for _, m := range all {
		if _, ok := applied[m.Version]; !ok {
			pending = append(pending, m)
		}
	}
	return pending
}

func appliedInReverseOrder(all []Migration, applied map[int]time.Time, steps int) []Migration {
	var toRevert []Migration
	for i := len(all) - 1; i >= 0; i-- {
		if _, ok := applied[all[i].Version]; ok {
			toRevert = append(toRevert, all[i])
			if len(toRevert) == steps {
				break
			}
		}
	}
	return toRevert
}

// =====================================================================
// CLI Helpers (used by main)
// =====================================================================

// newRunner opens a database connection using environment variables and
// returns a configured Runner pointed at the migrations directory.
func newRunner() (*Runner, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getenv("DB_HOST", "localhost"),
		getenv("DB_PORT", "5432"),
		getenv("DB_USER", "postgres"),
		getenv("DB_PASSWORD", "postgres"),
		getenv("DB_NAME", "mytodo_dev"),
		getenv("DB_SSLMODE", "disable"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return NewRunner(db, defaultMigrationDir), nil
}

// createMigration creates a new sequential .up.sql / .down.sql file pair
// in the migrations directory. Filenames follow the pattern:
//
//	NNNNNN_<name>.up.sql  /  NNNNNN_<name>.down.sql
//
// The sequence number is the current highest version + 1.
func createMigration(name string) error {
	dir := defaultMigrationDir
	if envDir := os.Getenv("MIGRATION_DIR"); envDir != "" {
		dir = envDir
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("reading migrations dir %s: %w", dir, err)
	}

	// Find the highest existing version number.
	maxVersion := 0
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		v, _, _, ok := parseMigrationFilename(e.Name())
		if ok && v > maxVersion {
			maxVersion = v
		}
	}

	next := maxVersion + 1
	// Sanitise name: lowercase, spaces → underscores.
	safeName := strings.ToLower(strings.ReplaceAll(strings.TrimSpace(name), " ", "_"))
	base := filepath.Join(dir, fmt.Sprintf("%06d_%s", next, safeName))

	upFile := base + ".up.sql"
	downFile := base + ".down.sql"

	upContent := fmt.Sprintf("-- Migration: %s\n-- Direction: UP\n-- Created: %s\n\n-- TODO: write your UP migration SQL here\n",
		safeName, time.Now().UTC().Format(time.RFC3339))
	downContent := fmt.Sprintf("-- Migration: %s\n-- Direction: DOWN\n-- Created: %s\n\n-- TODO: write your DOWN (rollback) SQL here\n",
		safeName, time.Now().UTC().Format(time.RFC3339))

	if err := os.WriteFile(upFile, []byte(upContent), 0644); err != nil {
		return fmt.Errorf("create up file: %w", err)
	}
	if err := os.WriteFile(downFile, []byte(downContent), 0644); err != nil {
		return fmt.Errorf("create down file: %w", err)
	}

	fmt.Printf("Created:\n  %s\n  %s\n", upFile, downFile)
	return nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
