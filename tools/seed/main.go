package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
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
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// Generous timeout for seed operations.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("ping db: %v\n\nMake sure the database is running and the connection vars are set:\n  DB_HOST DB_PORT DB_USER DB_PASSWORD DB_NAME", err)
	}

	dataDir := getenv("SEED_DATA_DIR", "tools/seed/data")

	if err := runSeed(ctx, db, dataDir); err != nil {
		log.Fatalf("seed failed: %v", err)
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
