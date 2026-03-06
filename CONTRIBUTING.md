# Contributing to MyTodo

First off, thank you for considering contributing to MyTodo! 🎉

The following is a set of guidelines for contributing to MyTodo. These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Architecture Guidelines](#architecture-guidelines)
- [Coding Standards](#coding-standards)
- [Testing Guidelines](#testing-guidelines)
- [Commit Messages](#commit-messages)
- [Pull Request Process](#pull-request-process)

---

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

---

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/mytodo.git
   cd mytodo
   ```
3. **Add upstream remote:**
   ```bash
   git remote add upstream https://github.com/ORIGINAL_OWNER/mytodo.git
   ```
4. **Set up development environment:**
   ```bash
   make setup
   ```

---

## Development Workflow

### 1. Create a Branch

Always create a new branch for your work:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/bug-description
```

### Branch Naming Convention

- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation changes
- `refactor/` - Code refactoring
- `test/` - Adding or updating tests
- `chore/` - Maintenance tasks

### 2. Make Your Changes

Follow the architecture and coding standards outlined below.

### 3. Run Tests

```bash
make test
```

### 4. Run Linters

```bash
make lint
```

### 5. Commit Your Changes

Follow our [commit message guidelines](#commit-messages).

### 6. Push to Your Fork

```bash
git push origin feature/your-feature-name
```

### 7. Create a Pull Request

Go to GitHub and create a pull request from your fork to the main repository.

---

## Architecture Guidelines

MyTodo follows **Clean Architecture** and **Domain-Driven Design** principles. Please maintain this structure:

### Layer Structure

```
Domain Layer
  ↓
Application Layer
  ↓
Infrastructure Layer
  ↓
Interface Layer
```

### When Adding a New Feature

Follow this order:

1. **Domain Layer** - Define entities, value objects, repository interfaces
   - `internal/{module}/domain/entity/`
   - `internal/{module}/domain/repository/`
   - `internal/{module}/domain/service/`

2. **Application Layer** - Define commands, queries, handlers
   - `internal/{module}/application/commands/`
   - `internal/{module}/application/queries/`
   - `internal/{module}/application/handlers/`

3. **Infrastructure Layer** - Implement repositories, external services
   - `internal/{module}/infrastructure/persistence/`
   - `internal/{module}/infrastructure/cache/`

4. **Interface Layer** - HTTP controllers, DTOs
   - `internal/{module}/interfaces/http/`
   - `internal/{module}/interfaces/dto/`

### Dependency Rules

- **Domain layer** must NOT depend on any other layer
- **Application layer** can depend only on Domain
- **Infrastructure layer** can depend on Domain and Application
- **Interface layer** can depend on all layers

---

## Coding Standards

### Go Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golangci-lint` for linting
- Maximum line length: 120 characters
- Use meaningful variable names

### Good Examples

```go
// Good: Clear, descriptive name
func CreateIssue(ctx context.Context, req CreateIssueRequest) (*Issue, error) {
    // Implementation
}

// Bad: Unclear, abbreviated
func CrtIss(c context.Context, r CrtIssReq) (*Iss, error) {
    // Implementation
}
```

### Error Handling

Always handle errors explicitly:

```go
// Good
result, err := someFunction()
if err != nil {
    return nil, fmt.Errorf("failed to execute some function: %w", err)
}

// Bad
result, _ := someFunction()
```

### Interfaces

Define interfaces where they're used, not where they're implemented:

```go
// Good: Interface in application layer
package application

type UserRepository interface {
    Create(ctx context.Context, user *domain.User) error
    FindByID(ctx context.Context, id string) (*domain.User, error)
}

// Implementation in infrastructure layer
package persistence

type PostgresUserRepository struct {
    // ...
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *domain.User) error {
    // ...
}
```

### Context

Always pass `context.Context` as the first parameter:

```go
func ProcessIssue(ctx context.Context, issueID string) error {
    // Implementation
}
```

---

## Testing Guidelines

### Test Coverage

- Aim for **>80% code coverage**
- All new features must include tests
- All bug fixes must include regression tests

### Test Types

#### 1. Unit Tests

Test individual functions/methods in isolation:

```go
func TestIssueService_CreateIssue(t *testing.T) {
    // Arrange
    mockRepo := &MockIssueRepository{}
    service := NewIssueService(mockRepo)
    
    // Act
    issue, err := service.CreateIssue(ctx, request)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, issue)
    assert.Equal(t, "TODO-1", issue.Key)
}
```

#### 2. Integration Tests

Test interactions with real external systems:

```go
// +build integration

func TestPostgresRepository_CreateIssue(t *testing.T) {
    // Use real database connection
    db := setupTestDB(t)
    defer db.Close()
    
    repo := NewPostgresIssueRepository(db)
    issue, err := repo.Create(ctx, testIssue)
    
    assert.NoError(t, err)
    assert.NotEmpty(t, issue.ID)
}
```

#### 3. E2E Tests

Test complete user flows:

```go
func TestCreateIssueFlow(t *testing.T) {
    // Start test server
    server := setupTestServer(t)
    defer server.Close()
    
    // Make HTTP request
    resp, err := http.Post(server.URL+"/api/v1/issues", ...)
    
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
```

### Test File Naming

- Unit tests: `*_test.go` in the same package
- Integration tests: `*_integration_test.go` with build tag
- E2E tests: In `tests/e2e/` directory

### Running Tests

```bash
# All tests
make test

# Unit tests only
make test-unit

# Integration tests
make test-integration

# E2E tests
make test-e2e

# With coverage
make coverage
```

---

## Commit Messages

We follow [Conventional Commits](https://www.conventionalcommits.org/).

### Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat` - New feature
- `fix` - Bug fix
- `docs` - Documentation changes
- `style` - Code style changes (formatting, etc.)
- `refactor` - Code refactoring
- `test` - Adding or updating tests
- `chore` - Maintenance tasks
- `perf` - Performance improvements

### Examples

```
feat(issues): add ability to assign multiple users to an issue

Implemented the feature to assign multiple users to a single issue.
Added new database table for issue_assignees and updated the API.

Closes #123
```

```
fix(auth): resolve JWT token expiration bug

Fixed a bug where JWT tokens were expiring prematurely due to
incorrect time calculation.

Fixes #456
```

### Rules

- Use present tense ("add feature" not "added feature")
- Keep subject line under 50 characters
- Use body to explain what and why, not how
- Reference issues and PRs in the footer

---

## Pull Request Process

### Before Submitting

1. ✅ All tests pass (`make test`)
2. ✅ Code is properly formatted (`make format`)
3. ✅ No linting errors (`make lint`)
4. ✅ Documentation is updated
5. ✅ Commit messages follow conventions
6. ✅ Branch is up to date with main

### PR Title

Follow the same convention as commit messages:

```
feat(issues): add bulk issue creation API
fix(auth): resolve session timeout issue
```

### PR Description Template

```markdown
## Description
Brief description of the changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Changes Made
- List of changes
- Another change
- One more change

## Testing
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] E2E tests added/updated
- [ ] Manual testing completed

## Checklist
- [ ] Code follows project style guidelines
- [ ] Self-review completed
- [ ] Comments added for complex code
- [ ] Documentation updated
- [ ] No new warnings generated
- [ ] Tests pass locally
- [ ] Dependent changes merged

## Screenshots (if applicable)
Add screenshots here

## Related Issues
Closes #issue_number
```

### Review Process

1. **At least one approval** required from a maintainer
2. **All CI checks must pass**
3. **No unresolved conversations**
4. **Branch must be up to date** with main

### After Approval

Maintainers will merge using **Squash and Merge** to keep the history clean.

---

## Code Review Guidelines

### As a Reviewer

- Be respectful and constructive
- Ask questions rather than make demands
- Acknowledge good practices
- Focus on the code, not the person
- Respond in a timely manner

### As an Author

- Be open to feedback
- Don't take criticism personally
- Respond to all comments
- Make requested changes or explain why not
- Test your changes after updates

---

## Database Migrations

### Creating a Migration

```bash
make migrate-create name=add_user_preferences_table
```

This creates two files:
- `YYYYMMDDHHMMSS_add_user_preferences_table.up.sql`
- `YYYYMMDDHHMMSS_add_user_preferences_table.down.sql`

### Migration Guidelines

- One logical change per migration
- Always include both UP and DOWN migrations
- Test both directions
- Use transactions where possible
- Never modify committed migrations

### Example Migration

```sql
-- up.sql
BEGIN;

CREATE TABLE user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    theme VARCHAR(20) DEFAULT 'light',
    language VARCHAR(10) DEFAULT 'en',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_preferences_user_id ON user_preferences(user_id);

COMMIT;

-- down.sql
BEGIN;

DROP TABLE IF EXISTS user_preferences;

COMMIT;
```

---

## Documentation

### Code Documentation

- Public functions must have GoDoc comments
- Complex algorithms need inline comments
- Include examples for non-obvious usage

```go
// CreateIssue creates a new issue in the specified project.
// It validates the input, generates a unique issue key, and persists
// the issue to the database.
//
// Example:
//   issue, err := service.CreateIssue(ctx, CreateIssueRequest{
//       ProjectID: "proj-123",
//       Title: "Bug in login flow",
//       Type: IssueTypeBug,
//   })
func CreateIssue(ctx context.Context, req CreateIssueRequest) (*Issue, error) {
    // Implementation
}
```

### API Documentation

- Update OpenAPI spec for new endpoints
- Include request/response examples
- Document all error responses

### Architecture Documentation

- Update architecture docs for significant changes
- Add ADRs (Architecture Decision Records) for major decisions
- Keep diagrams up to date

---

## Getting Help

- **Questions?** Open a [Discussion](https://github.com/yourrepo/mytodo/discussions)
- **Bug reports?** Open an [Issue](https://github.com/yourrepo/mytodo/issues)
- **Feature requests?** Open a [Discussion](https://github.com/yourrepo/mytodo/discussions)

---

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (MIT License).

---

**Thank you for contributing to MyTodo! 🚀**
