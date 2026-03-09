#!/bin/bash
# Migration script for running database migrations

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting database migrations...${NC}"

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | grep -v '^#' | xargs)
fi

# Database connection string
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_NAME=${DB_NAME:-mytodo}
DB_USER=${DB_USER:-postgres}
DB_PASSWORD=${DB_PASSWORD:-postgres}

DATABASE_URL="postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

echo -e "${YELLOW}Database: ${DB_NAME}${NC}"
echo -e "${YELLOW}Host: ${DB_HOST}:${DB_PORT}${NC}"

# Check if migrate tool is installed
if ! command -v migrate &> /dev/null; then
    echo -e "${RED}Error: migrate tool not found${NC}"
    echo "Install it with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
fi

# Run migrations
MIGRATION_PATH="file://apps/api/pkg/database/migrations"

case "$1" in
    up)
        echo -e "${GREEN}Running migrations up...${NC}"
        migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" up
        echo -e "${GREEN}Migrations completed successfully!${NC}"
        ;;
    down)
        echo -e "${YELLOW}Rolling back migrations...${NC}"
        migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" down $2
        echo -e "${GREEN}Rollback completed!${NC}"
        ;;
    drop)
        echo -e "${RED}Dropping all migrations...${NC}"
        read -p "Are you sure? This will drop all tables! (yes/no): " confirm
        if [ "$confirm" = "yes" ]; then
            migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" drop -f
            echo -e "${GREEN}All migrations dropped!${NC}"
        else
            echo -e "${YELLOW}Cancelled${NC}"
        fi
        ;;
    force)
        echo -e "${YELLOW}Forcing migration version to $2...${NC}"
        migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" force $2
        echo -e "${GREEN}Migration version forced!${NC}"
        ;;
    version)
        echo -e "${YELLOW}Current migration version:${NC}"
        migrate -path ${MIGRATION_PATH} -database "${DATABASE_URL}" version
        ;;
    create)
        if [ -z "$2" ]; then
            echo -e "${RED}Error: Migration name required${NC}"
            echo "Usage: $0 create <migration_name>"
            exit 1
        fi
        echo -e "${GREEN}Creating new migration: $2${NC}"
        migrate create -ext sql -dir apps/api/pkg/database/migrations -seq $2
        echo -e "${GREEN}Migration files created!${NC}"
        ;;
    *)
        echo "Usage: $0 {up|down|drop|force|version|create} [args]"
        echo ""
        echo "Commands:"
        echo "  up              - Run all pending migrations"
        echo "  down [N]        - Rollback last N migrations (default: 1)"
        echo "  drop            - Drop all migrations (requires confirmation)"
        echo "  force <version> - Force set migration version"
        echo "  version         - Show current migration version"
        echo "  create <name>   - Create a new migration"
        exit 1
        ;;
esac
