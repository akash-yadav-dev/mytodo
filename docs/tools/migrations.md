# Migrations Tool Documentation

**Location:** `tools/migrations/`  
**Purpose:** Manage database schema migrations

---

## Overview

The migrations tool is a CLI application for managing database schema changes in a controlled, versioned manner. It supports creating, applying, and rolling back migrations.

---

## Installation

```bash
# Run directly
go run tools/migrations/main.go [command]

# Or build binary
cd tools/migrations
go build -o migrate main.go
./migrate [command]
```

---

## Commands

### up

Apply all pending migrations.

```bash
go run tools/migrations/main.go up

# Apply specific number of migrations
go run tools/migrations/main.go up --steps=2
```

**Options:**
- `--steps`: Number of migrations to apply (default: all)
- `--dry-run`: Show what would be executed without applying

---

### down

Rollback migrations.

```bash
go run tools/migrations/main.go down

# Rollback specific number of migrations
go run tools/migrations/main.go down --steps=1
```

**Options:**
- `--steps`: Number of migrations to rollback (default: 1)
- `--force`: Skip confirmation prompt

---

### status

Show migration status.

```bash
go run tools/migrations/main.go status
```

**Output:**
```
Migration Status:
┌──────────────────────────┬────────────────────────────────┬─────────┬─────────────────────┐
│ Version                  │ Name                           │ Status  │ Applied At          │
├──────────────────────────┼────────────────────────────────┼─────────┼─────────────────────┤
│ 20260301120000          │ create_users_table             │ Applied │ 2026-03-01 12:05:23 │
│ 20260301130000          │ create_organizations_table     │ Applied │ 2026-03-01 13:02:15 │
│ 20260302140000          │ create_projects_table          │ Applied │ 2026-03-02 14:10:45 │
│ 20260303150000          │ create_issues_table            │ Pending │ -                   │
└──────────────────────────┴────────────────────────────────┴─────────┴─────────────────────┘

Pending: 1
Applied: 3
```

---

### create

Create a new migration file.

```bash
go run tools/migrations/main.go create add_custom_fields_table
```

**Output:**
```
Created migration files:
- apps/api/pkg/database/migrations/20260306120000_add_custom_fields_table.up.sql
- apps/api/pkg/database/migrations/20260306120000_add_custom_fields_table.down.sql
```

**Generated Template:**

`20260306120000_add_custom_fields_table.up.sql`:
```sql
-- Migration: add_custom_fields_table
-- Created: 2026-03-06 12:00:00

BEGIN;

-- Write your migration here
CREATE TABLE custom_fields (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMIT;
```

`20260306120000_add_custom_fields_table.down.sql`:
```sql
-- Rollback: add_custom_fields_table
-- Created: 2026-03-06 12:00:00

BEGIN;

-- Write your rollback here
DROP TABLE IF EXISTS custom_fields;

COMMIT;
```

---

### version

Get current migration version.

```bash
go run tools/migrations/main.go version
```

**Output:**
```
Current version: 20260302140000
Last migration: create_projects_table
```

---

### reset

Reset database (rollback all migrations).

```bash
go run tools/migrations/main.go reset --confirm
```

**⚠️ WARNING:** This will drop all tables and data!

---

### fresh

Reset database and rerun all migrations.

```bash
go run tools/migrations/main.go fresh --confirm
```

---

### validate

Validate migration files syntax.

```bash
go run tools/migrations/main.go validate
```

---

## Migration Files

### File Naming Convention

```
{timestamp}_{description}.{direction}.sql
```

Examples:
- `20260301120000_create_users_table.up.sql`
- `20260301120000_create_users_table.down.sql`

### File Structure

```sql
-- Migration: create_users_table
-- Created: 2026-03-01 12:00:00
-- Description: Create users table with authentication fields

BEGIN;

-- Migration statements here
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);

-- Add comments
COMMENT ON TABLE users IS 'User accounts';

COMMIT;
```

### Best Practices

1. **Always use transactions** (`BEGIN`/`COMMIT`)
2. **Write reversible migrations** (test down migrations)
3. **One logical change per migration**
4. **Add comments** for complex changes
5. **Test on dev database first**
6. **Never modify applied migrations** (create new ones instead)

---

## Configuration

### Environment Variables

```bash
# Database connection
DB_HOST=localhost
DB_PORT=5432
DB_USER=mytodo
DB_PASSWORD=mytodo_dev_password
DB_NAME=mytodo_dev
DB_SSL_MODE=disable

# Migration settings
MIGRATIONS_PATH=apps/api/pkg/database/migrations
MIGRATIONS_TABLE=schema_migrations
```

### Config File

Create `tools/migrations/config.yaml`:

```yaml
database:
  host: localhost
  port: 5432
  user: mytodo
  password: mytodo_dev_password
  database: mytodo_dev
  sslmode: disable

migrations:
  path: apps/api/pkg/database/migrations
  table: schema_migrations
  lock_timeout: 30s
```

---

## Advanced Usage

### Multiple Databases

```bash
# Run on different database
go run tools/migrations/main.go up --database=mytodo_test

# Or set environment variable
DB_NAME=mytodo_test go run tools/migrations/main.go up
```

---

### Dry Run

Preview what will be executed:

```bash
go run tools/migrations/main.go up --dry-run
```

**Output:**
```
[DRY RUN] Would execute:
- 20260303150000_create_issues_table.up.sql
- 20260304160000_create_comments_table.up.sql

No changes made to database.
```

---

### Parallel Migrations

Run migrations on multiple databases:

```bash
# migrations.sh
#!/bin/bash
databases=("mytodo_dev" "mytodo_staging" "mytodo_prod")

for db in "${databases[@]}"; do
    echo "Migrating $db..."
    DB_NAME=$db go run tools/migrations/main.go up
done
```

---

## Migration Examples

### Add Column

```sql
-- up.sql
BEGIN;
ALTER TABLE users ADD COLUMN phone_number VARCHAR(20);
CREATE INDEX idx_users_phone ON users(phone_number);
COMMIT;

-- down.sql
BEGIN;
DROP INDEX IF EXISTS idx_users_phone;
ALTER TABLE users DROP COLUMN IF EXISTS phone_number;
COMMIT;
```

---

### Add Table with Foreign Keys

```sql
-- up.sql
BEGIN;
CREATE TABLE project_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

CREATE INDEX idx_project_members_project ON project_members(project_id);
CREATE INDEX idx_project_members_user ON project_members(user_id);
COMMIT;

-- down.sql
BEGIN;
DROP TABLE IF EXISTS project_members;
COMMIT;
```

---

### Data Migration

```sql
-- up.sql
BEGIN;

-- Add new column
ALTER TABLE issues ADD COLUMN issue_number INTEGER;

-- Populate data
WITH numbered AS (
    SELECT 
        id,
        ROW_NUMBER() OVER (PARTITION BY project_id ORDER BY created_at) as num
    FROM issues
)
UPDATE issues i
SET issue_number = n.num
FROM numbered n
WHERE i.id = n.id;

-- Make column required
ALTER TABLE issues ALTER COLUMN issue_number SET NOT NULL;

COMMIT;
```

---

## Troubleshooting

### Issue: Migration stuck in "running" state

**Cause:** Previous migration crashed mid-execution

**Solution:**
```bash
# Force unlock
go run tools/migrations/main.go unlock --force

# Check status
go run tools/migrations/main.go status
```

---

### Issue: "migration already applied"

**Cause:** Migration version already exists in database

**Solution:**
```bash
# Check which migrations are applied
go run tools/migrations/main.go status

# If migration failed, mark as not applied
go run tools/migrations/main.go mark-unapplied 20260303150000
```

---

### Issue: Syntax error in migration

**Cause:** Invalid SQL in migration file

**Solution:**
```bash
# Validate all migrations
go run tools/migrations/main.go validate

# Test specific migration
psql -h localhost -U mytodo -d mytodo_dev -f migrations/20260303150000_create_issues.up.sql
```

---

## CI/CD Integration

### GitHub Actions

```yaml
name: Run Migrations

on:
  push:
    branches: [main, staging]

jobs:
  migrate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Setup Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      
      - name: Run migrations
        env:
          DB_HOST: ${{ secrets.DB_HOST }}
          DB_USER: ${{ secrets.DB_USER }}
          DB_PASSWORD: ${{ secrets.DB_PASSWORD }}
          DB_NAME: ${{ secrets.DB_NAME }}
        run: |
          go run tools/migrations/main.go up
```

---

## See Also

- [Database Schema](../database/schema.md)
- [Migration Guide](../database/migration-guide.md)
- [Seed Tool](seed.md)
