# Projects API Endpoints

**Base URL:** `/api/v1/projects`  
**Version:** 1.0  
**Authentication:** Required

---

## Create Project

Create a new project in an organization.

**Endpoint:** `POST /api/v1/projects`

**Request Body:**
```json
{
  "organization_id": "org-123",
  "name": "Marketing Campaign 2026",
  "key": "MKTG",
  "description": "Q1 2026 marketing initiatives",
  "visibility": "private",
  "lead_id": "user-123",
  "category": "marketing",
  "start_date": "2026-03-01",
  "end_date": "2026-06-30"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "project": {
      "id": "project-123",
      "organization_id": "org-123",
      "name": "Marketing Campaign 2026",
      "key": "MKTG",
      "description": "Q1 2026 marketing initiatives",
      "icon": "📢",
      "color": "#3B82F6",
      "visibility": "private",
      "lead": {
        "id": "user-123",
        "username": "johndoe",
        "full_name": "John Doe"
      },
      "category": "marketing",
      "start_date": "2026-03-01",
      "end_date": "2026-06-30",
      "status": "active",
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## List Projects

Get projects for organization.

**Endpoint:** `GET /api/v1/projects`

**Query Parameters:**
- `organization_id`: Filter by organization (required)
- `status`: Filter by status (active, archived, on_hold)
- `category`: Filter by category
- `search`: Search by name or key
- `limit`: Results limit (default: 50)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "projects": [
      {
        "id": "project-123",
        "name": "Marketing Campaign 2026",
        "key": "MKTG",
        "icon": "📢",
        "color": "#3B82F6",
        "lead": {
          "id": "user-123",
          "full_name": "John Doe",
          "avatar_url": "https://cdn.mytodo.com/avatars/user.jpg"
        },
        "member_count": 8,
        "issue_count": 45,
        "status": "active",
        "updated_at": "2026-03-06T09:00:00Z"
      }
    ],
    "pagination": {
      "total": 12,
      "limit": 50,
      "offset": 0
    }
  }
}
```

---

## Get Project

Get project details.

**Endpoint:** `GET /api/v1/projects/{project_id}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "project": {
      "id": "project-123",
      "organization_id": "org-123",
      "name": "Marketing Campaign 2026",
      "key": "MKTG",
      "description": "Q1 2026 marketing initiatives",
      "icon": "📢",
      "color": "#3B82F6",
      "visibility": "private",
      "lead": {
        "id": "user-123",
        "username": "johndoe",
        "full_name": "John Doe",
        "avatar_url": "https://cdn.mytodo.com/avatars/user.jpg"
      },
      "default_assignee": {
        "id": "user-124",
        "username": "janesmith",
        "full_name": "Jane Smith"
      },
      "category": "marketing",
      "start_date": "2026-03-01",
      "end_date": "2026-06-30",
      "status": "active",
      "settings": {
        "issue_types": ["task", "bug", "story"],
        "time_tracking": true,
        "require_description": false
      },
      "stats": {
        "total_issues": 45,
        "open_issues": 23,
        "in_progress_issues": 12,
        "done_issues": 10,
        "member_count": 8
      },
      "created_by": {
        "id": "user-123",
        "full_name": "John Doe"
      },
      "created_at": "2026-01-15T10:00:00Z",
      "updated_at": "2026-03-06T09:00:00Z"
    }
  }
}
```

---

## Update Project

Update project details.

**Endpoint:** `PATCH /api/v1/projects/{project_id}`

**Permissions:** Project lead or admin

**Request Body:**
```json
{
  "name": "Updated Project Name",
  "description": "New description",
  "icon": "🎯",
  "color": "#10B981",
  "lead_id": "user-125",
  "status": "active",
  "settings": {
    "time_tracking": false
  }
}
```

**Response:** `200 OK`

---

## Archive Project

Archive a project.

**Endpoint:** `POST /api/v1/projects/{project_id}/archive`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Project archived successfully"
}
```

---

## Restore Project

Restore an archived project.

**Endpoint:** `POST /api/v1/projects/{project_id}/restore`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Project restored successfully"
}
```

---

## Delete Project

Permanently delete a project.

**Endpoint:** `DELETE /api/v1/projects/{project_id}`

**Request Body:**
```json
{
  "confirmation": "DELETE"
}
```

**Response:** `200 OK`

---

## Project Members

### List Members

**Endpoint:** `GET /api/v1/projects/{project_id}/members`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "members": [
      {
        "id": "member-1",
        "user": {
          "id": "user-1",
          "username": "johndoe",
          "full_name": "John Doe",
          "avatar_url": "https://cdn.mytodo.com/avatars/user1.jpg"
        },
        "role": "lead",
        "permissions": ["manage_project", "manage_issues", "manage_members"],
        "joined_at": "2026-01-15T10:00:00Z"
      }
    ],
    "total": 8
  }
}
```

---

### Add Member

**Endpoint:** `POST /api/v1/projects/{project_id}/members`

**Request Body:**
```json
{
  "user_id": "user-125",
  "role": "developer"
}
```

**Response:** `201 Created`

---

### Update Member Role

**Endpoint:** `PATCH /api/v1/projects/{project_id}/members/{member_id}`

**Request Body:**
```json
{
  "role": "admin"
}
```

**Response:** `200 OK`

---

### Remove Member

**Endpoint:** `DELETE /api/v1/projects/{project_id}/members/{member_id}`

**Response:** `200 OK`

---

## Project Statistics

**Endpoint:** `GET /api/v1/projects/{project_id}/stats`

**Query Parameters:**
- `period`: Time period (week, month, quarter, year)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "stats": {
      "overview": {
        "total_issues": 45,
        "open_issues": 23,
        "in_progress": 12,
        "done": 10
      },
      "by_type": {
        "bug": 15,
        "task": 20,
        "story": 10
      },
      "by_priority": {
        "critical": 3,
        "high": 8,
        "medium": 20,
        "low": 14
      },
      "velocity": {
        "current_sprint": 18,
        "average": 15,
        "trend": "up"
      },
      "time_tracking": {
        "estimated_hours": 360,
        "logged_hours": 245,
        "remaining_hours": 115
      }
    }
  }
}
```

---

## Project Activity

**Endpoint:** `GET /api/v1/projects/{project_id}/activity`

**Query Parameters:**
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
        "user": {
          "id": "user-1",
          "full_name": "John Doe",
          "avatar_url": "https://cdn.mytodo.com/avatars/user1.jpg"
        },
        "entity": {
          "type": "issue",
          "id": "issue-123",
          "key": "MKTG-45",
          "title": "Create landing page"
        },
        "timestamp": "2026-03-06T10:00:00Z"
      }
    ],
    "pagination": {
      "next_cursor": "cursor-token",
      "has_more": true
    }
  }
}
```
