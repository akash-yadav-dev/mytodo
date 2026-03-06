# Seed Tool Documentation

**Location:** `tools/seed/`  
**Purpose:** Populate database with initial and test data

---

## Overview

The seed tool populates the database with initial data, test data, and sample datasets for development and testing environments.

---

## Installation

```bash
# Run directly
go run tools/seed/main.go [command]

# Or build binary
cd tools/seed
go build -o seeder main.go
./seeder [command]
```

---

## Commands

### run

Run all seeders or specific seeders.

```bash
# Run all seeders
go run tools/seed/main.go run

# Run specific seeders
go run tools/seed/main.go run --only=users,organizations

# Run with environment
go run tools/seed/main.go run --env=development
```

**Options:**
- `--env`: Environment (development, staging, production)
- `--only`: Comma-separated list of seeders to run
- `--except`: Comma-separated list of seeders to skip
- `--truncate`: Truncate tables before seeding

---

### list

List all available seeders.

```bash
go run tools/seed/main.go list
```

**Output:**
```
Available Seeders:
┌──────────────────────┬─────────────┬────────────────────────────────┐
│ Name                 │ Order       │ Description                     │
├──────────────────────┼─────────────┼────────────────────────────────┤
│ users                │ 1           │ Create admin and test users    │
│ organizations        │ 2           │ Create sample organizations    │
│ projects             │ 3           │ Create sample projects         │
│ issues               │ 4           │ Create sample issues and tasks │
│ comments             │ 5           │ Add comments to issues         │
└──────────────────────┴─────────────┴────────────────────────────────┘
```

---

### fresh

Drop all data and reseed.

```bash
go run tools/seed/main.go fresh --confirm
```

**⚠️ WARNING:** This will delete all data!

---

### factory

Generate data using factories.

```bash
# Generate 100 users
go run tools/seed/main.go factory --model=user --count=100

# Generate with custom attributes
go run tools/seed/main.go factory --model=issue --count=50 --attrs='{"status":"done"}'
```

---

## Seed Files

### Directory Structure

```
tools/seed/
├── main.go
├── seeder.go
├── data/              # Static seed data (JSON/YAML)
│   ├── users.json
│   ├── organizations.json
│   └── issue_types.json
└── factory/           # Data factories
    ├── user_factory.go
    ├── project_factory.go
    └── issue_factory.go
```

---

### Data Files

`tools/seed/data/users.json`:
```json
{
  "users": [
    {
      "email": "admin@mytodo.com",
      "username": "admin",
      "password": "Admin123!",
      "full_name": "System Administrator",
      "role": "admin"
    },
    {
      "email": "demo@mytodo.com",
      "username": "demo",
      "password": "Demo123!",
      "full_name": "Demo User",
      "role": "user"
    }
  ]
}
```

---

### Factory Files

`tools/seed/factory/user_factory.go`:
```go
package factory

import (
    "github.com/google/uuid"
    "github.com/brianvoe/gofakeit/v6"
)

type UserFactory struct{}

func (f *UserFactory) Create(count int, attrs map[string]interface{}) []map[string]interface{} {
    users := make([]map[string]interface{}, count)
    
    for i := 0; i < count; i++ {
        user := map[string]interface{}{
            "id":         uuid.New().String(),
            "email":      gofakeit.Email(),
            "username":   gofakeit.Username(),
            "full_name":  gofakeit.Name(),
            "avatar_url": gofakeit.ImageURL(200, 200),
            "timezone":   "America/New_York",
            "locale":     "en-US",
            "status":     "active",
        }
        
        // Override with custom attributes
        for key, value := range attrs {
            user[key] = value
        }
        
        users[i] = user
    }
    
    return users
}
```

---

## Seeder Examples

### Basic Seeder

```go
// tools/seed/seeders/users_seeder.go
package seeders

import (
    "context"
    "database/sql"
)

type UsersSeeder struct {
    db *sql.DB
}

func NewUsersSeeder(db *sql.DB) *UsersSeeder {
    return &UsersSeeder{db: db}
}

func (s *UsersSeeder) Run(ctx context.Context) error {
    users := []struct {
        Email    string
        Username string
        FullName string
    }{
        {"admin@mytodo.com", "admin", "Administrator"},
        {"demo@mytodo.com", "demo", "Demo User"},
    }
    
    for _, user := range users {
        _, err := s.db.ExecContext(ctx, `
            INSERT INTO users (email, username, full_name, email_verified)
            VALUES ($1, $2, $3, true)
            ON CONFLICT (email) DO NOTHING
        `, user.Email, user.Username, user.FullName)
        
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

---

### Seeder with Relations

```go
// tools/seed/seeders/projects_seeder.go
package seeders

type ProjectsSeeder struct {
    db *sql.DB
}

func (s *ProjectsSeeder) Run(ctx context.Context) error {
    // Get organization ID
    var orgID string
    err := s.db.QueryRowContext(ctx, `
        SELECT id FROM organizations LIMIT 1
    `).Scan(&orgID)
    if err != nil {
        return err
    }
    
    // Get user ID for lead
    var userID string
    err = s.db.QueryRowContext(ctx, `
        SELECT id FROM users WHERE email = 'admin@mytodo.com'
    `).Scan(&userID)
    if err != nil {
        return err
    }
    
    // Create projects
    projects := []struct {
        Name string
        Key  string
    }{
        {"Backend Development", "BACK"},
        {"Frontend Development", "FRONT"},
        {"Mobile App", "MOBILE"},
    }
    
    for _, proj := range projects {
        _, err := s.db.ExecContext(ctx, `
            INSERT INTO projects (organization_id, name, key, lead_id, status)
            VALUES ($1, $2, $3, $4, 'active')
            ON CONFLICT (organization_id, key) DO NOTHING
        `, orgID, proj.Name, proj.Key, userID)
        
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

---

## Configuration

### Environment Variables

```bash
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=mytodo
DB_PASSWORD=mytodo_dev_password
DB_NAME=mytodo_dev

# Seed settings
SEED_ENV=development
SEED_DATA_PATH=tools/seed/data
```

---

### Config File

`tools/seed/config.yaml`:
```yaml
database:
  host: localhost
  port: 5432
  user: mytodo
  password: mytodo_dev_password
  database: mytodo_dev

seed:
  environment: development
  data_path: tools/seed/data
  truncate_before: false
  
  # Seeder execution order
  order:
    - users
    - organizations
    - organization_members
    - projects
    - project_members
    - issue_types
    - issue_statuses
    - issue_priorities
    - issues
    - comments
    - attachments
```

---

## Environment-Specific Seeding

### Development

```bash
# Full dataset with realistic data
go run tools/seed/main.go run --env=development
```

Creates:
- 100+ users
- 10+ organizations
- 50+ projects
- 500+ issues
- 1000+ comments

---

### Staging

```bash
# Production-like data without PII
go run tools/seed/main.go run --env=staging
```

Creates:
- Anonymized user data
- Real project structures
- Sample issues
- No sensitive information

---

### Production

```bash
# Minimal essential data only
go run tools/seed/main.go run --env=production
```

Creates:
- System admin user
- Default issue types
- Default statuses
- Default priorities

---

## Factory Usage

### User Factory

```bash
# Generate 50 users
go run tools/seed/main.go factory --model=user --count=50

# Generate admins
go run tools/seed/main.go factory --model=user --count=5 --attrs='{"status":"admin"}'
```

---

### Issue Factory

```bash
# Generate 100 issues
go run tools/seed/main.go factory --model=issue --count=100

# Generate high-priority bugs
go run tools/seed/main.go factory --model=issue --count=20 \
    --attrs='{"issue_type":"bug","priority":"high"}'
```

---

### Project Factory

```bash
# Generate 10 projects
go run tools/seed/main.go factory --model=project --count=10
```

---

## Advanced Usage

### Custom Seeder

Create a new seeder:

```bash
# 1. Create seeder file
cat > tools/seed/seeders/custom_seeder.go << 'EOF'
package seeders

import (
    "context"
    "database/sql"
)

type CustomSeeder struct {
    db *sql.DB
}

func NewCustomSeeder(db *sql.DB) *CustomSeeder {
    return &CustomSeeder{db: db}
}

func (s *CustomSeeder) Run(ctx context.Context) error {
    // Your seeding logic here
    return nil
}
EOF

# 2. Register in main.go
# Add to seeders list

# 3. Run
go run tools/seed/main.go run --only=custom
```

---

### Conditional Seeding

```go
func (s *IssuesSeeder) Run(ctx context.Context) error {
    // Only seed if table is empty
    var count int
    err := s.db.QueryRowContext(ctx, `
        SELECT COUNT(*) FROM issues
    `).Scan(&count)
    if err != nil {
        return err
    }
    
    if count > 0 {
        log.Println("Issues already seeded, skipping...")
        return nil
    }
    
    // Seed logic...
    return nil
}
```

---

### Transaction Support

```go
func (s *ProjectsSeeder) Run(ctx context.Context) error {
    tx, err := s.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // Seeding operations...
    _, err = tx.ExecContext(ctx, `INSERT INTO projects ...`)
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

---

## Makefile Integration

```makefile
# Makefile

.PHONY: seed seed-fresh seed-dev seed-staging

seed:
	go run tools/seed/main.go run

seed-fresh:
	go run tools/seed/main.go fresh --confirm

seed-dev:
	go run tools/seed/main.go run --env=development

seed-staging:
	go run tools/seed/main.go run --env=staging
```

Usage:
```bash
make seed
make seed-dev
make seed-fresh
```

---

## Troubleshooting

### Issue: Foreign key constraint violation

**Cause:** Seeders running in wrong order

**Solution:** Check seeder order in config
```yaml
seed:
  order:
    - users         # Must run before projects
    - organizations # Must run before projects
    - projects      # Can now reference users and orgs
```

---

### Issue: Duplicate key errors

**Cause:** Running seeder multiple times

**Solution:** Use `ON CONFLICT` or check existence
```sql
INSERT INTO users (email, username, full_name)
VALUES ($1, $2, $3)
ON CONFLICT (email) DO NOTHING;
```

---

### Issue: Out of memory with large datasets

**Cause:** Loading too much data at once

**Solution:** Use batch processing
```go
batchSize := 1000
for i := 0; i < totalCount; i += batchSize {
    // Process batch
}
```

---

## CI/CD Integration

### Docker Compose

```yaml
# docker-compose.yml
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_DB: mytodo_dev
      POSTGRES_USER: mytodo
      POSTGRES_PASSWORD: mytodo_dev_password
  
  seeder:
    build: .
    depends_on:
      - postgres
    command: go run tools/seed/main.go run --env=development
    environment:
      DB_HOST: postgres
```

---

### GitHub Actions

```yaml
- name: Seed database
  run: |
    go run tools/migrations/main.go up
    go run tools/seed/main.go run --env=test
```

---

## See Also

- [Migrations Tool](migrations.md)
- [Database Schema](../database/schema.md)
- [Development Guide](../CONTRIBUTING.md)
