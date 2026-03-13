package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "up":
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := r.Up(); err != nil {
			log.Fatalf("migrate up: %v", err)
		}

	case "down":
		steps := 1
		if len(os.Args) >= 3 {
			n, err := strconv.Atoi(os.Args[2])
			if err != nil || n < 1 {
				log.Fatalf("invalid step count: %s", os.Args[2])
			}
			steps = n
		}
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := r.Down(steps); err != nil {
			log.Fatalf("migrate down: %v", err)
		}

	case "create":
		if len(os.Args) < 3 {
			log.Fatal("usage: migrate create <migration_name>")
		}
		if err := createMigration(os.Args[2]); err != nil {
			log.Fatalf("migrate create: %v", err)
		}

	case "status":
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := r.Status(); err != nil {
			log.Fatalf("migrate status: %v", err)
		}

	case "version":
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		v, err := r.Version()
		if err != nil {
			log.Fatalf("migrate version: %v", err)
		}
		fmt.Printf("Current migration version: %d\n", v)

	case "force":
		if len(os.Args) < 3 {
			log.Fatal("usage: migrate force <version>")
		}
		v, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("invalid version: %s", os.Args[2])
		}
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := r.Force(v); err != nil {
			log.Fatalf("migrate force: %v", err)
		}

	case "drop":
		fmt.Print("This will roll back ALL migrations. Type 'yes' to confirm: ")
		var confirm string
		fmt.Scanln(&confirm)
		if confirm != "yes" {
			fmt.Println("Cancelled.")
			return
		}
		r, err := newRunner()
		if err != nil {
			log.Fatalf("connect: %v", err)
		}
		if err := r.Drop(); err != nil {
			log.Fatalf("migrate drop: %v", err)
		}

	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Print(`Usage: migrate <command> [args]

Commands:
  up              Apply all pending migrations
  down [N]        Roll back the last N migrations (default 1)
  create <name>   Create a new migration file pair
  status          Show applied/pending migration status
  version         Print the current migration version
  force <N>       Force-set the schema version (recover dirty state)
  drop            Roll back ALL migrations (destructive)

Environment:
  DB_HOST         Postgres host      (default: localhost)
  DB_PORT         Postgres port      (default: 5432)
  DB_USER         Postgres user      (default: postgres)
  DB_PASSWORD     Postgres password  (default: postgres)
  DB_NAME         Postgres database  (default: mytodo_dev)
  MIGRATION_DIR   Path to migrations (default: apps/api/pkg/database/migrations)
`)
}
