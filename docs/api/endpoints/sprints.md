# Sprints API Endpoints

**Base URL:** `/api/v1/sprints`  
**Version:** 1.0  
**Authentication:** Required

---

## Create Sprint

Create a new sprint.

**Endpoint:** `POST /api/v1/sprints`

**Request Body:**
```json
{
  "project_id": "project-123",
  "board_id": "board-123",
  "name": "Sprint 5",
  "goal": "Complete authentication refactor and fix critical bugs",
  "start_date": "2026-03-10T00:00:00Z",
  "end_date": "2026-03-24T00:00:00Z"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "sprint": {
      "id": "sprint-5",
      "project_id": "project-123",
      "board_id": "board-123",
      "name": "Sprint 5",
      "goal": "Complete authentication refactor and fix critical bugs",
      "sprint_number": 5,
      "status": "planned",
      "start_date": "2026-03-10T00:00:00Z",
      "end_date": "2026-03-24T00:00:00Z",
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## List Sprints

Get sprints for a project.

**Endpoint:** `GET /api/v1/sprints`

**Query Parameters:**
- `project_id`: Filter by project (required)
- `board_id`: Filter by board
- `status`: Filter by status (planned, active, completed, cancelled)

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "sprints": [
      {
        "id": "sprint-5",
        "name": "Sprint 5",
        "sprint_number": 5,
        "goal": "Complete authentication refactor",
        "status": "active",
        "start_date": "2026-03-10T00:00:00Z",
        "end_date": "2026-03-24T00:00:00Z",
        "total_points": 45,
        "completed_points": 28,
        "total_issues": 18,
        "completed_issues": 11,
        "days_remaining": 4
      }
    ],
    "total": 12
  }
}
```

---

## Get Sprint

Get sprint details with issues.

**Endpoint:** `GET /api/v1/sprints/{sprint_id}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "sprint": {
      "id": "sprint-5",
      "project": {
        "id": "project-123",
        "key": "PROJ",
        "name": "My Project"
      },
      "name": "Sprint 5",
      "goal": "Complete authentication refactor and fix critical bugs",
      "sprint_number": 5,
      "status": "active",
      "start_date": "2026-03-10T00:00:00Z",
      "end_date": "2026-03-24T00:00:00Z",
      "completed_at": null,
      "total_points": 45,
      "completed_points": 28,
      "total_issues": 18,
      "completed_issues": 11,
      "issues": [
        {
          "id": "issue-1",
          "issue_key": "PROJ-123",
          "title": "Implement JWT refresh",
          "status": {
            "name": "Done",
            "category": "done"
          },
          "story_points": 8,
          "assignee": {
            "id": "user-1",
            "full_name": "John Doe"
          }
        }
      ],
      "created_by": {
        "id": "user-123",
        "full_name": "John Doe"
      },
      "created_at": "2026-02-28T10:00:00Z",
      "updated_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## Update Sprint

**Endpoint:** `PATCH /api/v1/sprints/{sprint_id}`

**Request Body:**
```json
{
  "name": "Updated Sprint Name",
  "goal": "Updated sprint goal",
  "start_date": "2026-03-11T00:00:00Z",
  "end_date": "2026-03-25T00:00:00Z"
}
```

**Response:** `200 OK`

---

## Start Sprint

Activate a planned sprint.

**Endpoint:** `POST /api/v1/sprints/{sprint_id}/start`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Sprint started successfully"
}
```

---

## Complete Sprint

Complete an active sprint.

**Endpoint:** `POST /api/v1/sprints/{sprint_id}/complete`

**Request Body:**
```json
{
  "move_incomplete_to": "sprint-6"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "summary": {
      "completed_issues": 11,
      "incomplete_issues": 7,
      "completed_points": 28,
      "incomplete_points": 17,
      "moved_to_sprint": "Sprint 6"
    }
  },
  "message": "Sprint completed successfully"
}
```

---

## Cancel Sprint

Cancel a sprint.

**Endpoint:** `POST /api/v1/sprints/{sprint_id}/cancel`

**Response:** `200 OK`

---

## Delete Sprint

**Endpoint:** `DELETE /api/v1/sprints/{sprint_id}`

**Response:** `200 OK`

---

## Add Issues to Sprint

Add issues to sprint.

**Endpoint:** `POST /api/v1/sprints/{sprint_id}/issues`

**Request Body:**
```json
{
  "issue_ids": ["issue-1", "issue-2", "issue-3"]
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "3 issues added to sprint"
}
```

---

## Remove Issue from Sprint

**Endpoint:** `DELETE /api/v1/sprints/{sprint_id}/issues/{issue_id}`

**Response:** `200 OK`

---

## Sprint Report

Get sprint burndown and velocity report.

**Endpoint:** `GET /api/v1/sprints/{sprint_id}/report`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "report": {
      "sprint": {
        "id": "sprint-5",
        "name": "Sprint 5",
        "status": "active"
      },
      "summary": {
        "total_points": 45,
        "completed_points": 28,
        "remaining_points": 17,
        "completion_rate": 62.2,
        "total_issues": 18,
        "completed_issues": 11,
        "days_elapsed": 10,
        "days_remaining": 4
      },
      "burndown": [
        {
          "date": "2026-03-10",
          "ideal_remaining": 45,
          "actual_remaining": 45
        },
        {
          "date": "2026-03-11",
          "ideal_remaining": 42,
          "actual_remaining": 40
        }
      ],
      "velocity": {
        "current_sprint": 28,
        "average_last_5": 32,
        "trend": "down"
      }
    }
  }
}
```
