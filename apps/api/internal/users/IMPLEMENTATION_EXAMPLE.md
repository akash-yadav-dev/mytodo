# User GetByID Endpoint - Complete Implementation Guide

This document shows the complete flow of the `GET /api/v1/users/:id` endpoint, demonstrating how all layers work together in Clean Architecture.

## 🏗️ Architecture Flow

```
HTTP Request
    ↓
[1] Routes (user_routes.go)
    ↓
[2] Controller (user_controller.go)
    ↓
[3] Query Handler (user_query_handler.go)
    ↓
[4] Repository (user_repository.go)
    ↓
[5] Database
```

## 📝 Complete Implementation

### 1. Route Definition (`user_routes.go`)

**Location:** `apps/api/internal/users/interfaces/http/user_routes.go`

```go
func RegisterUserRoutes(router *gin.RouterGroup, controller *UserController, jwtService *security.JWTService) {
    users := router.Group("/users")
    users.Use(middleware.AuthMiddleware(jwtService), loggingMiddleware)

    users.GET("/:id", controller.GetUser)  // ← Route for GET /api/v1/users/:id
}
```

**What happens:**
- Registers the route pattern `/users/:id`
- Applies authentication middleware
- Maps to `controller.GetUser` method

---

### 2. HTTP Controller (`user_controller.go`)

**Location:** `apps/api/internal/users/interfaces/http/user_controller.go`

```go
// GetUser retrieves a user by ID.
// GET /api/v1/users/:id
func (c *UserController) GetUser(ctx *gin.Context) {
    // Step 1: Extract user ID from URL parameter
    userID := ctx.Param("id")

    // Step 2: Create and execute query
    query := queries.GetUserByIDQuery{
        UserID:         userID,
        IncludeProfile: false,
    }

    user, err := c.queryHandler.HandleGetByID(query)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{
            "error":   "User not found",
            "message": err.Error(),
        })
        return
    }

    // Step 3: Return successful response
    ctx.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    user,
    })
}
```

**Responsibilities:**
- Extract parameters from HTTP request
- Create query object
- Call application handler
- Format HTTP response
- Handle HTTP-specific concerns (status codes, headers)

---

### 3. Query Handler (`user_query_handler.go`)

**Location:** `apps/api/internal/users/application/handlers/user_query_handler.go`

```go
// HandleGetByID retrieves a user by their unique identifier.
func (h *UserQueryHandler) HandleGetByID(query queries.GetUserByIDQuery) (*dto.UserDTO, error) {
    // Step 1: Validate query
    if err := query.Validate(); err != nil {
        return nil, err
    }

    // Step 2: Fetch user from repository
    user, err := h.userRepo.FindByID(query.UserID)
    if err != nil {
        return nil, errors.New("user not found")
    }

    // Step 3: Optionally fetch profile (if needed)
    // if query.IncludeProfile {
    //     profile, err := h.profileRepo.FindByUserID(query.UserID)
    //     ...
    // }

    // Step 4: Convert entity to DTO
    userDTO := dto.ToUserDTO(user)

    // Step 5: Return result
    return userDTO, nil
}
```

**Responsibilities:**
- Validate input
- Orchestrate business logic
- Call repository for data access
- Transform domain entities to DTOs
- Apply business rules

---

### 4. Query Object (`get_user_by_id_query.go`)

**Location:** `apps/api/internal/users/application/queries/get_user_by_id_query.go`

```go
type GetUserByIDQuery struct {
    UserID         string `json:"user_id" validate:"required"`
    IncludeProfile bool   `json:"include_profile"`
}

func (q GetUserByIDQuery) Validate() error {
    if q.UserID == "" {
        return errors.New("user_id is required")
    }
    return nil
}
```

**Purpose:**
- Encapsulates query parameters
- Provides validation logic
- Makes intent explicit (CQRS pattern)

---

### 5. Repository Interface (`user_repository.go`)

**Location:** `apps/api/internal/users/domain/repository/user_repository.go`

```go
type UserRepository interface {
    FindByID(id string) (*entity.User, error)
    FindByEmail(email string) (*entity.User, error)
    // ... other methods
}
```

**Purpose:**
- Defines data access contract
- Database-agnostic interface
- Enables dependency injection and testing

---

### 6. DTO and Mapper (`user_dto.go`)

**Location:** `apps/api/internal/users/interfaces/dto/user_dto.go`

```go
type UserDTO struct {
    ID          string    `json:"id"`
    Email       string    `json:"email"`
    Username    string    `json:"username"`
    DisplayName string    `json:"display_name"`
    Avatar      string    `json:"avatar,omitempty"`
    Bio         string    `json:"bio,omitempty"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

func ToUserDTO(user *entity.User) *UserDTO {
    if user == nil {
        return nil
    }
    return &UserDTO{
        ID:          user.ID,
        Email:       user.Email,
        Username:    user.Username,
        DisplayName: user.DisplayName,
        Avatar:      user.Avatar,
        Bio:         user.Bio,
        Status:      user.Status,
        CreatedAt:   user.CreatedAt,
        UpdatedAt:   user.UpdatedAt,
    }
}
```

**Purpose:**
- Defines API contract
- Separates external representation from domain model
- Enables API versioning without changing domain

---

## 🔄 Complete Request Flow Example

### Request
```http
GET /api/v1/users/550e8400-e29b-41d4-a716-446655440000
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

### Step-by-Step Execution

1. **Gin Router** receives request at `/api/v1/users/:id`

2. **AuthMiddleware** validates JWT token
   - Extracts user_id from token
   - Sets in context: `ctx.Set("user_id", "authenticated-user-id")`

3. **Route Handler** calls `controller.GetUser(ctx)`

4. **Controller** extracts parameter:
   ```go
   userID := ctx.Param("id") // "550e8400-e29b-41d4-a716-446655440000"
   ```

5. **Controller** creates query:
   ```go
   query := queries.GetUserByIDQuery{
       UserID: "550e8400-e29b-41d4-a716-446655440000",
   }
   ```

6. **Handler** validates query:
   ```go
   query.Validate() // Returns nil if valid
   ```

7. **Handler** fetches from repository:
   ```go
   user, err := h.userRepo.FindByID("550e8400...")
   // Returns: &entity.User{ID: "550e...", Email: "john@example.com", ...}
   ```

8. **Handler** converts to DTO:
   ```go
   userDTO := dto.ToUserDTO(user)
   // Returns: &dto.UserDTO{ID: "550e...", Email: "john@example.com", ...}
   ```

9. **Controller** sends JSON response:
   ```json
   {
     "success": true,
     "data": {
       "id": "550e8400-e29b-41d4-a716-446655440000",
       "email": "john@example.com",
       "username": "johndoe",
       "display_name": "John Doe",
       "status": "active",
       "created_at": "2024-01-15T10:30:00Z",
       "updated_at": "2024-03-01T14:20:00Z"
     }
   }
   ```

### Response
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "success": true,
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "email": "john@example.com",
    "username": "johndoe",
    "display_name": "John Doe",
    "avatar": "https://cdn.example.com/avatars/john.jpg",
    "status": "active",
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-03-01T14:20:00Z"
  }
}
```

---

## 📂 File Locations Reference

```
apps/api/internal/users/
├── application/
│   ├── handlers/
│   │   └── user_query_handler.go      [3] Query Handler
│   └── queries/
│       └── get_user_by_id_query.go     [4] Query Object
├── domain/
│   ├── entity/
│   │   └── user.go                     [6] Domain Entity
│   └── repository/
│       └── user_repository.go          [5] Repository Interface
└── interfaces/
    ├── dto/
    │   └── user_dto.go                 [7] DTO & Mapper
    └── http/
        ├── user_controller.go          [2] HTTP Controller
        └── user_routes.go              [1] Route Registration
```

---

## 🎯 Key Principles Applied

### 1. **Separation of Concerns**
- Routes handle routing logic
- Controllers handle HTTP concerns
- Handlers handle business logic
- Repositories handle data access

### 2. **Dependency Inversion**
- Handlers depend on repository interfaces (not implementations)
- Enables testing and swapping implementations

### 3. **Single Responsibility**
- Each layer has one clear purpose
- Easy to modify without affecting others

### 4. **CQRS Pattern**
- Queries (reads) separated from Commands (writes)
- Different optimization strategies for each

### 5. **DTO Pattern**
- API contracts independent of domain model
- Safe to expose externally
- Enables API versioning

---

## 🧪 Testing Strategy

### Unit Test: Handler
```go
func TestHandleGetByID(t *testing.T) {
    // Mock repository
    mockRepo := &MockUserRepository{
        users: map[string]*entity.User{
            "123": {ID: "123", Email: "test@example.com"},
        },
    }
    
    handler := NewUserQueryHandler(mockRepo, nil)
    
    // Test successful retrieval
    query := queries.GetUserByIDQuery{UserID: "123"}
    result, err := handler.HandleGetByID(query)
    
    assert.NoError(t, err)
    assert.Equal(t, "test@example.com", result.Email)
}
```

### Integration Test: Controller
```go
func TestGetUserEndpoint(t *testing.T) {
    router := setupTestRouter()
    
    req := httptest.NewRequest("GET", "/api/v1/users/123", nil)
    req.Header.Set("Authorization", "Bearer "+testToken)
    w := httptest.NewRecorder()
    
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "test@example.com")
}
```

---

## ✅ Checklist for New Endpoints

When creating new endpoints, follow this pattern:

- [ ] Create query/command object in `application/queries` or `application/commands`
- [ ] Add validation to query/command
- [ ] Implement handler method in `application/handlers`
- [ ] Add repository method to interface (if needed)
- [ ] Create/update DTO in `interfaces/dto`
- [ ] Add controller method in `interfaces/http`
- [ ] Register route in `interfaces/http/routes.go`
- [ ] Write unit tests for handler
- [ ] Write integration tests for controller
- [ ] Update API documentation

---

## 🚀 Next Steps

To implement other endpoints, follow this same pattern:

1. **GET /api/v1/users** (List) - Already stubbed in handler
2. **PUT /api/v1/users/:id** (Update) - Create UpdateUserCommand
3. **DELETE /api/v1/users/:id** (Delete) - Create DeleteUserCommand
4. **GET /api/v1/users/search** (Search) - Use SearchUsersQuery

Each follows the same architectural flow!
