# Users API Endpoints

**Base URL:** `/api/v1/users`  
**Version:** 1.0  
**Authentication:** Required (Bearer token)

---

## Get Current User

Get authenticated user profile.

**Endpoint:** `GET /api/v1/users/me`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "username": "johndoe",
      "full_name": "John Doe",
      "avatar_url": "https://cdn.mytodo.com/avatars/user.jpg",
      "timezone": "America/New_York",
      "locale": "en-US",
      "email_verified": true,
      "phone_number": "+1234567890",
      "phone_verified": false,
      "two_factor_enabled": true,
      "status": "active",
      "preferences": {
        "theme": "dark",
        "notifications_email": true,
        "notifications_push": true,
        "language": "en"
      },
      "created_at": "2026-01-15T10:00:00Z",
      "updated_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## Update Current User

Update authenticated user profile.

**Endpoint:** `PATCH /api/v1/users/me`

**Request Body:**
```json
{
  "full_name": "John Smith",
  "timezone": "Europe/London",
  "locale": "en-GB",
  "phone_number": "+441234567890",
  "preferences": {
    "theme": "light",
    "notifications_email": false
  }
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "user": { /* updated user object */ }
  },
  "message": "Profile updated successfully"
}
```

---

## Upload Avatar

Upload user avatar image.

**Endpoint:** `POST /api/v1/users/me/avatar`

**Request:** `multipart/form-data`
```
avatar: [File]
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "avatar_url": "https://cdn.mytodo.com/avatars/user-new.jpg"
  },
  "message": "Avatar uploaded successfully"
}
```

---

## Delete Avatar

Remove user avatar.

**Endpoint:** `DELETE /api/v1/users/me/avatar`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Avatar deleted successfully"
}
```

---

## Get User by ID

Get public user profile.

**Endpoint:** `GET /api/v1/users/{user_id}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "username": "johndoe",
      "full_name": "John Doe",
      "avatar_url": "https://cdn.mytodo.com/avatars/user.jpg",
      "status": "active"
    }
  }
}
```

---

## Search Users

Search for users in organization/project.

**Endpoint:** `GET /api/v1/users/search`

**Query Parameters:**
- `q`: Search query (username, name, email)
- `organization_id`: Filter by organization
- `project_id`: Filter by project
- `limit`: Results limit (default: 20)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "users": [
      {
        "id": "user-1",
        "username": "johndoe",
        "full_name": "John Doe",
        "avatar_url": "https://cdn.mytodo.com/avatars/user1.jpg",
        "email": "j***@example.com"
      }
    ],
    "pagination": {
      "total": 45,
      "limit": 20,
      "offset": 0,
      "has_more": true
    }
  }
}
```

---

## Get User Activity

Get user activity timeline.

**Endpoint:** `GET /api/v1/users/{user_id}/activity`

**Query Parameters:**
- `type`: Activity type filter (issue, comment, worklog)
- `limit`: Results limit (default: 50)
- `before`: Cursor for pagination

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "activities": [
      {
        "id": "activity-1",
        "type": "issue_created",
        "entity": {
          "id": "issue-1",
          "key": "PROJ-123",
          "title": "Fix login bug"
        },
        "timestamp": "2026-03-06T10:00:00Z"
      },
      {
        "id": "activity-2",
        "type": "comment_added",
        "entity": {
          "id": "comment-1",
          "issue_key": "PROJ-122",
          "content": "Working on this now"
        },
        "timestamp": "2026-03-06T09:30:00Z"
      }
    ],
    "pagination": {
      "next_cursor": "cursor-token",
      "has_more": true
    }
  }
}
```

---

## Get User Statistics

Get user statistics and metrics.

**Endpoint:** `GET /api/v1/users/{user_id}/stats`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "stats": {
      "issues_created": 125,
      "issues_assigned": 89,
      "issues_completed": 67,
      "comments_made": 234,
      "time_logged_minutes": 15600,
      "contributions_this_month": 45,
      "velocity_avg_points": 13
    }
  }
}
```

---

## Delete Account

Permanently delete user account.

**Endpoint:** `DELETE /api/v1/users/me`

**Request Body:**
```json
{
  "password": "CurrentPassword123!",
  "confirmation": "DELETE"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Account deleted successfully"
}
```
