package main

// defaultMigrationDir is the canonical path to the SQL migration files,
// relative to the workspace root. The Runner will use this when no
// MIGRATION_DIR env var is set.
const defaultMigrationDir = "apps/api/pkg/database/migrations"
