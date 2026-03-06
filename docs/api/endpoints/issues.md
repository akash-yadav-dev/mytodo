# Issues API Endpoints

**Base URL:** `/api/v1/issues`  
**Version:** 1.0  
**Authentication:** Required

---

## Create Issue

Create a new issue.

**Endpoint:** `POST /api/v1/issues`

**Request Body:**
```json
{
  "project_id": "project-123",
  "title": "Fix login bug on mobile",
  "description": "Users are unable to log in on iOS devices",
  "issue_type_id": "type-bug",
  "status_id": "status-backlog",
  "priority_id": "priority-high",
  "assignee_id": "user-123",
  "labels": ["mobile", "authentication"],
  "story_points": 5,
  "due_date": "2026-03-15T00:00:00Z",
  "custom_fields": {
    "browser": "Safari",
    "os_version": "iOS 17"
  }
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "issue": {
      "id": "issue-123",
      "issue_key": "PROJ-124",
      "issue_number": 124,
      "project": {
        "id": "project-123",
        "key": "PROJ",
        "name": "My Project"
      },
      "title": "Fix login bug on mobile",
      "description": "Users are unable to log in on iOS devices",
      "issue_type": {
        "id": "type-bug",
        "name": "Bug",
        "icon": "🐛",
        "color": "#EF4444"
      },
      "status": {
        "id": "status-backlog",
        "name": "Backlog",
        "category": "todo"
      },
      "priority": {
        "id": "priority-high",
        "name": "High",
        "level": 2
      },
      "reporter": {
        "id": "user-122",
        "username": "reporter",
        "full_name": "Reporter Name",
        "avatar_url": "url"
      },
      "assignee": {
        "id": "user-123",
        "username": "johndoe",
        "full_name": "John Doe",
        "avatar_url": "url"
      },
      "labels": ["mobile", "authentication"],
      "story_points": 5,
      "estimate_minutes": null,
      "time_spent_minutes": 0,
      "due_date": "2026-03-15T00:00:00Z",
      "created_at": "2026-03-06T10:00:00Z",
      "updated_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## List Issues

Get issues with filtering and pagination.

**Endpoint:** `GET /api/v1/issues`

**Query Parameters:**
- `project_id`: Filter by project (optional)
- `organization_id`: Filter by organization (optional)
- `assignee_id`: Filter by assignee
- `reporter_id`: Filter by reporter
- `status_id`: Filter by status
- `priority_id`: Filter by priority
- `issue_type_id`: Filter by issue type
- `sprint_id`: Filter by sprint
- `epic_id`: Filter by epic
- `labels`: Comma-separated labels
- `search`: Full-text search query
- `sort`: Sort field (created_at, updated_at, priority)
- `order`: Sort order (asc, desc)
- `limit`: Results limit (default: 50)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "issues": [
      {
        "id": "issue-123",
        "issue_key": "PROJ-124",
        "title": "Fix login bug on mobile",
        "issue_type": {
          "name": "Bug",
          "icon": "🐛",
          "color": "#EF4444"
        },
        "status": {
          "name": "In Progress",
          "category": "in_progress"
        },
        "priority": {
          "name": "High",
          "level": 2
        },
        "assignee": {
          "id": "user-123",
          "full_name": "John Doe",
          "avatar_url": "url"
        },
        "labels": ["mobile", "authentication"],
        "story_points": 5,
        "due_date": "2026-03-15T00:00:00Z",
        "created_at": "2026-03-06T10:00:00Z",
        "updated_at": "2026-03-06T11:30:00Z"
      }
    ],
    "pagination": {
      "total": 234,
      "limit": 50,
      "offset": 0,
      "has_more": true
    }
  }
}
```

---

## Get Issue

Get detailed issue information.

**Endpoint:** `GET /api/v1/issues/{issue_id}`

or

**Endpoint:** `GET /api/v1/issues/by-key/{issue_key}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "issue": {
      "id": "issue-123",
      "issue_key": "PROJ-124",
      "issue_number": 124,
      "project": {
        "id": "project-123",
        "key": "PROJ",
        "name": "My Project",
        "icon": "📱"
      },
      "title": "Fix login bug on mobile",
      "description": "Users are unable to log in on iOS devices. Error occurs when...",
      "description_html": "<p>Users are unable to log in...</p>",
      "issue_type": {
        "id": "type-bug",
        "name": "Bug",
        "icon": "🐛",
        "color": "#EF4444"
      },
      "status": {
        "id": "status-inprogress",
        "name": "In Progress",
        "category": "in_progress",
        "color": "#F59E0B"
      },
      "priority": {
        "id": "priority-high",
        "name": "High",
        "level": 2,
        "icon": "🟠"
      },
      "reporter": {
        "id": "user-122",
        "username": "reporter",
        "full_name": "Reporter Name",
        "avatar_url": "url"
      },
      "assignee": {
        "id": "user-123",
        "username": "johndoe",
        "full_name": "John Doe",
        "avatar_url": "url"
      },
      "parent_issue": null,
      "epic": {
        "id": "epic-1",
        "key": "PROJ-1",
        "title": "Mobile App Improvements"
      },
      "sprint": {
        "id": "sprint-5",
        "name": "Sprint 5",
        "status": "active"
      },
      "labels": ["mobile", "authentication"],
      "story_points": 5,
      "estimate_minutes": 480,
      "time_spent_minutes": 180,
      "remaining_estimate_minutes": 300,
      "due_date": "2026-03-15T00:00:00Z",
      "start_date": "2026-03-06T10:00:00Z",
      "resolved_at": null,
      "custom_fields": {
        "browser": "Safari",
        "os_version": "iOS 17",
        "severity": "Critical"
      },
      "comment_count": 5,
      "attachment_count": 2,
      "subtask_count": 3,
      "watcher_count": 4,
      "vote_count": 2,
      "watchers": [
        {
          "id": "user-125",
          "full_name": "Watcher Name",
          "avatar_url": "url"
        }
      ],
      "links": [
        {
          "id": "link-1",
          "type": "blocks",
          "target_issue": {
            "id": "issue-125",
            "key": "PROJ-125",
            "title": "Related issue"
          }
        }
      ],
      "created_by": {
        "id": "user-122",
        "full_name": "Creator Name"
      },
      "updated_by": {
        "id": "user-123",
        "full_name": "John Doe"
      },
      "created_at": "2026-03-06T10:00:00Z",
      "updated_at": "2026-03-06T11:30:00Z"
    }
  }
}
```

---

## Update Issue

Update issue fields.

**Endpoint:** `PATCH /api/v1/issues/{issue_id}`

**Request Body:**
```json
{
  "title": "Updated title",
  "description": "Updated description",
  "status_id": "status-done",
  "priority_id": "priority-medium",
  "assignee_id": "user-124",
  "labels": ["mobile", "authentication", "urgent"],
  "story_points": 8,
  "estimate_minutes": 600,
  "due_date": "2026-03-20T00:00:00Z",
  "custom_fields": {
    "environment": "production"
  }
}
```

**Response:** `200 OK`

---

## Transition Issue

Move issue to a different status.

**Endpoint:** `POST /api/v1/issues/{issue_id}/transitions`

**Request Body:**
```json
{
  "status_id": "status-inreview",
  "comment": "Ready for review",
  "resolution": "fixed"
}
```

**Response:** `200 OK`

---

## Delete Issue

Soft delete an issue.

**Endpoint:** `DELETE /api/v1/issues/{issue_id}`

**Response:** `200 OK`

---

## Restore Issue

Restore a deleted issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/restore`

**Response:** `200 OK`

---

## Assign Issue

Assign issue to a user.

**Endpoint:** `POST /api/v1/issues/{issue_id}/assign`

**Request Body:**
```json
{
  "assignee_id": "user-123"
}
```

**Response:** `200 OK`

---

## Unassign Issue

Remove assignee from issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/unassign`

**Response:** `200 OK`

---

## Watch Issue

Add current user as watcher.

**Endpoint:** `POST /api/v1/issues/{issue_id}/watch`

**Response:** `200 OK`

---

## Unwatch Issue

Remove current user from watchers.

**Endpoint:** `DELETE /api/v1/issues/{issue_id}/watch`

**Response:** `200 OK`

---

## Vote Issue

Upvote an issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/vote`

**Response:** `200 OK`

---

## Unvote Issue

Remove vote from issue.

**Endpoint:** `DELETE /api/v1/issues/{issue_id}/vote`

**Response:** `200 OK`

---

## Link Issues

Create a link between two issues.

**Endpoint:** `POST /api/v1/issues/{issue_id}/links`

**Request Body:**
```json
{
  "target_issue_id": "issue-125",
  "link_type": "blocks"
}
```

**Link Types:**
- `blocks` / `is_blocked_by`
- `relates_to`
- `duplicates` / `is_duplicated_by`
- `causes` / `is_caused_by`
- `parent_of` / `child_of`

**Response:** `201 Created`

---

## Delete Issue Link

Remove link between issues.

**Endpoint:** `DELETE /api/v1/issues/{issue_id}/links/{link_id}`

**Response:** `200 OK`

---

## Get Issue History

Get complete change history.

**Endpoint:** `GET /api/v1/issues/{issue_id}/history`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "history": [
      {
        "id": "history-1",
        "field_name": "status",
        "old_value": "To Do",
        "new_value": "In Progress",
        "changed_by": {
          "id": "user-123",
          "full_name": "John Doe"
        },
        "changed_at": "2026-03-06T11:00:00Z"
      }
    ],
    "total": 12
  }
}
```

---

## Get Issue Activity

Get combined activity feed (comments, history, worklogs).

**Endpoint:** `GET /api/v1/issues/{issue_id}/activity`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "activities": [
      {
        "type": "comment",
        "data": {
          "id": "comment-1",
          "content": "Working on this now",
          "user": {
            "id": "user-123",
            "full_name": "John Doe"
          }
        },
        "timestamp": "2026-03-06T11:30:00Z"
      },
      {
        "type": "field_change",
        "data": {
          "field": "status",
          "old_value": "To Do",
          "new_value": "In Progress",
          "user": {
            "id": "user-123",
            "full_name": "John Doe"
          }
        },
        "timestamp": "2026-03-06T11:00:00Z"
      }
    ]
  }
}
```

---

## Bulk Update Issues

Update multiple issues at once.

**Endpoint:** `PATCH /api/v1/issues/bulk`

**Request Body:**
```json
{
  "issue_ids": ["issue-1", "issue-2", "issue-3"],
  "updates": {
    "status_id": "status-done",
    "labels": ["reviewed"]
  }
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "updated_count": 3,
    "failed_count": 0
  }
}
```

---

## Export Issues

Export issues to CSV/JSON.

**Endpoint:** `GET /api/v1/issues/export`

**Query Parameters:**
- `format`: Export format (csv, json, excel)
- `project_id`: Project filter
- `status_id`: Status filter
- (All other list filters apply)

**Response:** `200 OK`
- Content-Type: `text/csv` or `application/json`

---

## Advanced Search

Advanced issue search with JQL-like query.

**Endpoint:** `POST /api/v1/issues/search`

**Request Body:**
```json
{
  "query": {
    "and": [
      {
        "field": "project_id",
        "operator": "equals",
        "value": "project-123"
      },
      {
        "field": "status.category",
        "operator": "in",
        "value": ["todo", "in_progress"]
      },
      {
        "or": [
          {
            "field": "priority.level",
            "operator": "lte",
            "value": 2
          },
          {
            "field": "labels",
            "operator": "contains",
            "value": "urgent"
          }
        ]
      }
    ]
  },
  "sort": [
    { "field": "priority.level", "order": "asc" },
    { "field": "created_at", "order": "desc" }
  ],
  "limit": 50,
  "offset": 0
}
```

**Response:** `200 OK` (Same format as List Issues)
