# Boards API Endpoints

**Base URL:** `/api/v1/boards`  
**Version:** 1.0  
**Authentication:** Required

---

## Create Board

Create a new board (Kanban or Scrum).

**Endpoint:** `POST /api/v1/boards`

**Request Body:**
```json
{
  "project_id": "project-123",
  "name": "Development Board",
  "description": "Main development workflow",
  "board_type": "scrum",
  "filter_config": {
    "issue_types": ["story", "bug", "task"],
    "assignees": ["user-123", "user-124"]
  }
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "board": {
      "id": "board-123",
      "project_id": "project-123",
      "name": "Development Board",
      "description": "Main development workflow",
      "board_type": "scrum",
      "is_default": false,
      "filter_config": {},
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## Get Board

Get board with columns and Issues.

**Endpoint:** `GET /api/v1/boards/{board_id}`

**Query Parameters:**
- `sprint_id`: Filter by sprint (for Scrum boards)
- `include_backlog`: Include backlog issues (default: false)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "board": {
      "id": "board-123",
      "name": "Development Board",
      "board_type": "scrum",
      "columns": [
        {
          "id": "column-1",
          "name": "To Do",
          "status": {
            "id": "status-todo",
            "name": "To Do"
          },
          "position": 0,
          "wip_limit": null,
          "issues": [
            {
              "id": "issue-1",
              "issue_key": "PROJ-1",
              "title": "Issue title",
              "issue_type": { "name": "Story", "icon": "📖" },
              "priority": { "name": "High", "level": 2 },
              "assignee": {
                "id": "user-1",
                "full_name": "John Doe",
                "avatar_url": "url"
              },
              "story_points": 5,
              "board_position": 0
            }
          ],
          "issue_count": 12
        }
      ]
    }
  }
}
```

---

## Update Board

**Endpoint:** `PATCH /api/v1/boards/{board_id}`

**Request Body:**
```json
{
  "name": "Updated Board Name",
  "description": "Updated description"
}
```

**Response:** `200 OK`

---

## Delete Board

**Endpoint:** `DELETE /api/v1/boards/{board_id}`

**Response:** `200 OK`

---

## Board Columns

### Create Column

**Endpoint:** `POST /api/v1/boards/{board_id}/columns`

**Request Body:**
```json
{
  "name": "In Review",
  "status_id": "status-review",
  "position": 2,
  "wip_limit": 5
}
```

**Response:** `201 Created`

---

### Update Column

**Endpoint:** `PATCH /api/v1/boards/{board_id}/columns/{column_id}`

**Request Body:**
```json
{
  "name": "Updated Column",
  "wip_limit": 10,
  "position": 3
}
```

**Response:** `200 OK`

---

### Delete Column

**Endpoint:** `DELETE /api/v1/boards/{board_id}/columns/{column_id}`

**Response:** `200 OK`

---

### Reorder Columns

**Endpoint:** `POST /api/v1/boards/{board_id}/columns/reorder`

**Request Body:**
```json
{
  "column_order": [
    { "id": "column-1", "position": 0 },
    { "id": "column-2", "position": 1 },
    { "id": "column-3", "position": 2 }
  ]
}
```

**Response:** `200 OK`

---

## Move Issue on Board

Move issue to different column/position.

**Endpoint:** `POST /api/v1/boards/{board_id}/issues/{issue_id}/move`

**Request Body:**
```json
{
  "column_id": "column-2",
  "position": 3
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Issue moved successfully"
}
```
