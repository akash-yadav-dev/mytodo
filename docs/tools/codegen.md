# Code Generator Documentation

**Location:** `tools/codegen/`  
**Purpose:** Generate boilerplate code for models, APIs, and more

---

## Overview

The code generator automates creation of repetitive code patterns following Clean Architecture and DDD principles used in the project.

---

## Installation

```bash
# Run directly
go run tools/codegen/main.go [command]

# Or build binary
cd tools/codegen
go build -o codegen main.go
./codegen [command]
```

---

## Commands

### model

Generate domain models from database schema.

```bash
# Generate model for a table
go run tools/codegen/main.go model --table=users

# Generate all models
go run tools/codegen/main.go model --all

# Generate with custom package
go run tools/codegen/main.go model --table=issues --package=issues
```

**Options:**
- `--table`: Table name
- `--all`: Generate for all tables
- `--package`: Target package name
- `--output`: Output directory
- `--force`: Overwrite existing files

---

### api

Generate API handlers for a domain.

```bash
# Generate complete API for a domain
go run tools/codegen/main.go api --domain=issues

# Generate only specific handler
go run tools/codegen/main.go api --domain=issues --handler=create
```

**Generates:**
- HTTP handlers
- Request/Response DTOs
- Input validation
- Route registration
- Tests

---

### service

Generate service layer code.

```bash
go run tools/codegen/main.go service --domain=issues
```

**Generates:**
- Service interface
- Service implementation
- Business logic stubs
- Unit tests

---

### repository

Generate repository layer code.

```bash
go run tools/codegen/main.go repository --domain=issues --table=issues
```

**Generates:**
- Repository interface
- PostgreSQL implementation
- Query methods (CRUD + custom)
- Integration tests

---

### crud

Generate complete CRUD operations.

```bash
go run tools/codegen/main.go crud --domain=labels --table=labels
```

**Generates:**
- Domain entity
- Repository interface and implementation
- Service layer
- HTTP handlers
- DTOs
- Tests
- Routes

---

### migration

Generate migration from model changes.

```bash
go run tools/codegen/main.go migration --name=add_archived_field --table=projects
```

---

### test

Generate test files for existing code.

```bash
# Generate tests for service
go run tools/codegen/main.go test --type=service --file=apps/api/internal/issues/domain/service/issue_service.go

# Generate tests for handler
go run tools/codegen/main.go test --type=handler --file=apps/api/internal/issues/interfaces/http/issue_handler.go
```

---

## Usage Examples

### Generate New Domain

Create a complete new domain with all layers:

```bash
# 1. Generate database migration
go run tools/codegen/main.go migration --name=create_labels_table

# Edit the migration file, then apply
go run tools/migrations/main.go up

# 2. Generate complete CRUD
go run tools/codegen/main.go crud --domain=labels --table=labels

# 3. Review and customize generated code
# Files created:
# - apps/api/internal/labels/domain/entity/label.go
# - apps/api/internal/labels/domain/repository/label_repository.go
# - apps/api/internal/labels/domain/service/label_service.go
# - apps/api/internal/labels/application/commands/
# - apps/api/internal/labels/application/queries/
# - apps/api/internal/labels/infrastructure/persistence/
# - apps/api/internal/labels/interfaces/http/
# - apps/api/internal/labels/interfaces/dto/
```

---

### Generate Model from Database

```bash
# Generate model for issues table
go run tools/codegen/main.go model --table=issues --output=apps/api/internal/issues/domain/entity
```

**Generated** (`issue.go`):
```go
package entity

import (
    "time"
    "github.com/google/uuid"
)

// Issue represents an issue/task entity
type Issue struct {
    ID             uuid.UUID  `json:"id" db:"id"`
    ProjectID      uuid.UUID  `json:"project_id" db:"project_id"`
    IssueNumber    int        `json:"issue_number" db:"issue_number"`
    IssueKey       string     `json:"issue_key" db:"issue_key"`
    Title          string     `json:"title" db:"title"`
    Description    *string    `json:"description" db:"description"`
    IssueTypeID    uuid.UUID  `json:"issue_type_id" db:"issue_type_id"`
    StatusID       uuid.UUID  `json:"status_id" db:"status_id"`
    PriorityID     *uuid.UUID `json:"priority_id" db:"priority_id"`
    AssigneeID     *uuid.UUID `json:"assignee_id" db:"assignee_id"`
    ReporterID     uuid.UUID  `json:"reporter_id" db:"reporter_id"`
    StoryPoints    *int       `json:"story_points" db:"story_points"`
    DueDate        *time.Time `json:"due_date" db:"due_date"`
    CreatedAt      time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
    DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}

// TableName returns the table name
func (Issue) TableName() string {
    return "issues"
}
```

---

### Generate API Handlers

```bash
go run tools/codegen/main.go api --domain=issues
```

**Generated** (`issue_handler.go`):
```go
package http

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "mytodo/apps/api/internal/issues/domain/service"
    "mytodo/apps/api/internal/issues/interfaces/dto"
)

type IssueHandler struct {
    service service.IssueService
}

func NewIssueHandler(service service.IssueService) *IssueHandler {
    return &IssueHandler{
        service: service,
    }
}

// CreateIssue handles POST /api/v1/issues
func (h *IssueHandler) CreateIssue(c *gin.Context) {
    var req dto.CreateIssueRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    issue, err := h.service.CreateIssue(c.Request.Context(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, dto.NewIssueResponse(issue))
}

// GetIssue handles GET /api/v1/issues/:id
func (h *IssueHandler) GetIssue(c *gin.Context) {
    id := c.Param("id")
    
    issue, err := h.service.GetIssueByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
        return
    }
    
    c.JSON(http.StatusOK, dto.NewIssueResponse(issue))
}

// Additional CRUD methods...

// RegisterRoutes registers all routes for issues
func (h *IssueHandler) RegisterRoutes(r *gin.RouterGroup) {
    issues := r.Group("/issues")
    {
        issues.POST("", h.CreateIssue)
        issues.GET("/:id", h.GetIssue)
        issues.PATCH("/:id", h.UpdateIssue)
        issues.DELETE("/:id", h.DeleteIssue)
        issues.GET("", h.ListIssues)
    }
}
```

---

### Generate Repository

```bash
go run tools/codegen/main.go repository --domain=issues --table=issues
```

**Generated** (`issue_repository.go`):
```go
package repository

import (
    "context"
    "mytodo/apps/api/internal/issues/domain/entity"
)

// IssueRepository defines the interface for issue data access
type IssueRepository interface {
    Create(ctx context.Context, issue *entity.Issue) error
    GetByID(ctx context.Context, id string) (*entity.Issue, error)
    Update(ctx context.Context, issue *entity.Issue) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, filter *IssueFilter) ([]*entity.Issue, error)
    Count(ctx context.Context, filter *IssueFilter) (int, error)
}

type IssueFilter struct {
    ProjectID  *string
    StatusID   *string
    AssigneeID *string
    Search     string
    Limit      int
    Offset     int
}
```

**Implementation** (`postgres_issue_repository.go`):
```go
package persistence

import (
    "context"
    "database/sql"
    "mytodo/apps/api/internal/issues/domain/entity"
    "mytodo/apps/api/internal/issues/domain/repository"
)

type postgresIssueRepository struct {
    db *sql.DB
}

func NewPostgresIssueRepository(db *sql.DB) repository.IssueRepository {
    return &postgresIssueRepository{db: db}
}

func (r *postgresIssueRepository) Create(ctx context.Context, issue *entity.Issue) error {
    query := `
        INSERT INTO issues (
            project_id, title, description, issue_type_id,
            status_id, priority_id, assignee_id, reporter_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, issue_number, issue_key, created_at
    `
    
    return r.db.QueryRowContext(
        ctx, query,
        issue.ProjectID, issue.Title, issue.Description,
        issue.IssueTypeID, issue.StatusID, issue.PriorityID,
        issue.AssigneeID, issue.ReporterID,
    ).Scan(&issue.ID, &issue.IssueNumber, &issue.IssueKey, &issue.CreatedAt)
}

// Additional methods...
```

---

## Templates

The code generator uses customizable templates:

```
tools/codegen/templates/
├── model.tmpl
├── handler.tmpl
├── service.tmpl
├── repository.tmpl
├── dto.tmpl
├── test.tmpl
└── migration.tmpl
```

### Custom Template

Create custom template:

```go
// tools/codegen/templates/custom.tmpl
package {{.Package}}

// {{.Name}} is a custom generated struct
type {{.Name}} struct {
    {{range .Fields}}
    {{.Name}} {{.Type}} `json:"{{.JSONTag}}"`
    {{end}}
}
```

Use custom template:
```bash
go run tools/codegen/main.go generate \
    --template=tools/codegen/templates/custom.tmpl \
    --output=output.go \
    --data='{"Package":"mypackage","Name":"MyStruct"}'
```

---

## Configuration

`tools/codegen/config.yaml`:
```yaml
output:
  base_path: apps/api/internal
  
templates:
  path: tools/codegen/templates
  
naming:
  style: snake_case  # or camelCase, PascalCase
  
generation:
  add_timestamps: true
  add_soft_delete: true
  add_validation: true
  add_documentation: true
  
database:
  host: localhost
  port: 5432
  user: mytodo
  password: mytodo_dev_password
  database: mytodo_dev
```

---

## Advanced Features

### Custom Fields

Add custom fields to generated models:

```bash
go run tools/codegen/main.go model --table=issues \
    --add-field="CustomField string json:\"custom_field\""
```

---

### Relationship Inference

Automatically detect and generate relationships:

```bash
go run tools/codegen/main.go model --table=issues --with-relations
```

**Generated:**
```go
type Issue struct {
    // ... fields ...
    
    // Relations
    Project  *Project  `json:"project,omitempty"`
    Assignee *User     `json:"assignee,omitempty"`
    Reporter *User     `json:"reporter,omitempty"`
    Comments []Comment `json:"comments,omitempty"`
}
```

---

### API Documentation

Generate OpenAPI/Swagger docs:

```bash
go run tools/codegen/main.go openapi --domain=issues --output=docs/api/openapi.yaml
```

---

## Best Practices

1. **Review generated code** before committing
2. **Customize templates** for your conventions
3. **Use --dry-run** to preview changes
4. **Version control templates**
5. **Keep generated and manual code separate**
6. **Add // Code generated DO NOT EDIT comments**

---

## Troubleshooting

### Issue: Template parsing error

**Cause:** Invalid template syntax

**Solution:**
```bash
# Validate templates
go run tools/codegen/main.go validate-templates
```

---

### Issue: Can't connect to database

**Cause:** Database not running or wrong credentials

**Solution:** Check environment variables and connection

---

### Issue: Generated code has compile errors

**Cause:** Type mapping issue

**Solution:** Customize type mappings in config

---

## See Also

- [Project Structure](../../plan/folder%20strcuture%20and%20project%20architecture.md)
- [Clean Architecture Guide](../architecture/clean-architecture.md)
- [Domain-Driven Design](../architecture/ddd.md)
