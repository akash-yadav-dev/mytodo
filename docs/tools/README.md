# Tools Documentation

This directory contains documentation for all development tools in the MyTodo project.

---

## Available Tools

### 1. [Migrations Tool](migrations.md)
Database migration management CLI for creating, running, and rolling back database schema changes.

**Location:** `tools/migrations/`

---

### 2. [Seed Tool](seed.md)
Database seeding tool for populating the database with initial and test data.

**Location:** `tools/seed/`

---

### 3. [Code Generator](codegen.md)
Code generation tool for creating boilerplate code, models, API handlers, and more.

**Location:** `tools/codegen/`

---

## Quick Start

### Prerequisites
- Go 1.21 or higher
- PostgreSQL 15+
- Access to development database

### Installation

```bash
# Navigate to project root
cd mytodo

# Run any tool directly with go run
go run tools/migrations/main.go --help
go run tools/seed/main.go --help
go run tools/codegen/main.go --help
```

---

## Common Workflows

### Setting up a new database

```bash
# 1. Create and run migrations
go run tools/migrations/main.go up

# 2. Seed initial data
go run tools/seed/main.go --env=development

# 3. Verify database
go run tools/migrations/main.go status
```

### Creating a new feature

```bash
# 1. Create migration for new table
go run tools/migrations/main.go create add_feature_table

# 2. Generate model code
go run tools/codegen/main.go model --table=features

# 3. Generate API handlers
go run tools/codegen/main.go api --domain=features
```

---

## Environment Variables

All tools support these environment variables:

```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=mytodo
DB_PASSWORD=mytodo_dev_password
DB_NAME=mytodo_dev
DB_SSL_MODE=disable
```

Alternatively, use a `.env` file in the project root.

---

## Best Practices

1. **Always test migrations** on a development database first
2. **Back up data** before running destructive migrations
3. **Use descriptive names** for migrations and seed files
4. **Version control everything** including generated code
5. **Document custom changes** to generated code

---

## Troubleshooting

### Common Issues

**Issue:** "connection refused"
- **Solution:** Ensure PostgreSQL is running and credentials are correct

**Issue:** "migration already applied"
- **Solution:** Check migration status with `status` command

**Issue:** "generated file already exists"
- **Solution:** Use `--force` flag to overwrite existing files

---

## See Also

- [Database Schema](../database/schema.md)
- [Migration Guide](../database/migration-guide.md)
- [Development Setup](../CONTRIBUTING.md)
