# Workflows, Automation & Integrations API

**Base URL:** `/api/v1`  
**Version:** 1.0  
**Authentication:** Required

---

## Workflows

### Create Workflow

**Endpoint:** `POST /api/v1/workflows`

**Request Body:**
```json
{
  "organization_id": "org-123",
  "project_id": "project-123",
  "name": "Development Workflow",
  "description": "Standard development workflow",
  "is_default": false
}
```

**Response:** `201 Created`

---

### List Workflows

**Endpoint:** `GET /api/v1/workflows`

**Query Parameters:**
- `organization_id`: Filter by organization
- `project_id`: Filter by project

---

## Workflow Transitions

### Create Transition

**Endpoint:** `POST /api/v1/workflows/{workflow_id}/transitions`

**Request Body:**
```json
{
  "name": "Start Progress",
  "from_status_id": "status-todo",
  "to_status_id": "status-inprogress",
  "required_fields": ["assignee"],
  "conditions": {
    "must_be_assigned": true
  },
  "post_functions": [
    {
      "type": "send_notification",
      "config": {
        "notify": "assignee"
      }
    }
  ]
}
```

**Response:** `201 Created`

---

## Automation Rules

### Create Automation Rule

**Endpoint:** `POST /api/v1/automation/rules`

**Request Body:**
```json
{
  "organization_id": "org-123",
  "project_id": "project-123",
  "name": "Auto-assign high priority bugs",
  "description": "Automatically assign high priority bugs to team lead",
  "trigger_type": "issue_created",
  "trigger_config": {},
  "conditions": [
    {
      "field": "issue_type",
      "operator": "equals",
      "value": "bug"
    },
    {
      "field": "priority.level",
      "operator": "lte",
      "value": 2
    }
  ],
  "actions": [
    {
      "type": "assign_issue",
      "config": {
        "assign_to": "user-123"
      }
    },
    {
      "type": "add_comment",
      "config": {
        "content": "Auto-assigned due to high priority"
      }
    },
    {
      "type": "send_notification",
      "config": {
        "notify_users": ["user-123"],
        "message": "New high priority bug assigned"
      }
    }
  ],
  "is_active": true
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "rule": {
      "id": "rule-123",
      "name": "Auto-assign high priority bugs",
      "trigger_type": "issue_created",
      "is_active": true,
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

### List Automation Rules

**Endpoint:** `GET /api/v1/automation/rules`

**Query Parameters:**
- `organization_id`: Required
- `project_id`: Optional
- `trigger_type`: Filter by trigger
- `is_active`: Filter by status

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "rules": [
      {
        "id": "rule-123",
        "name": "Auto-assign high priority bugs",
        "trigger_type": "issue_created",
        "is_active": true,
        "execution_count": 45,
        "last_executed_at": "2026-03-06T09:30:00Z"
      }
    ],
    "total": 5
  }
}
```

---

### Update Automation Rule

**Endpoint:** `PATCH /api/v1/automation/rules/{rule_id}`

---

### Toggle Automation Rule

**Endpoint:** `POST /api/v1/automation/rules/{rule_id}/toggle`

**Request Body:**
```json
{
  "is_active": false
}
```

**Response:** `200 OK`

---

### Delete Automation Rule

**Endpoint:** `DELETE /api/v1/automation/rules/{rule_id}`

**Response:** `200 OK`

---

### Test Automation Rule

Test rule without executing actions.

**Endpoint:** `POST /api/v1/automation/rules/{rule_id}/test`

**Request Body:**
```json
{
  "issue_id": "issue-123"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "matched": true,
    "conditions_met": true,
    "actions_that_would_execute": [
      "assign_issue",
      "add_comment",
      "send_notification"
    ]
  }
}
```

---

## Integrations

### Create Integration

**Endpoint:** `POST /api/v1/integrations`

**Request Body:**
```json
{
  "organization_id": "org-123",
  "project_id": "project-123",
  "integration_type": "github",
  "name": "GitHub - Main Repo",
  "config": {
    "repository": "myorg/myrepo",
    "sync_issues": true,
    "sync_prs": true
  },
  "credentials": {
    "access_token": "ghp_encrypted_token"
  }
}
```

**Response:** `201 Created`

---

### List Integrations

**Endpoint:** `GET /api/v1/integrations`

**Query Parameters:**
- `organization_id`: Required
- `project_id`: Optional
- `integration_type`: Filter by type

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "integrations": [
      {
        "id": "integration-123",
        "integration_type": "github",
        "name": "GitHub - Main Repo",
        "is_active": true,
        "last_sync_at": "2026-03-06T09:00:00Z",
        "sync_status": "success"
      }
    ],
    "total": 3
  }
}
```

---

### Get Integration

**Endpoint:** `GET /api/v1/integrations/{integration_id}`

**Response:** `200 OK`

---

### Update Integration

**Endpoint:** `PATCH /api/v1/integrations/{integration_id}`

**Response:** `200 OK`

---

### Delete Integration

**Endpoint:** `DELETE /api/v1/integrations/{integration_id}`

**Response:** `200 OK`

---

### Test Integration

**Endpoint:** `POST /api/v1/integrations/{integration_id}/test`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "connection": "success",
    "message": "Successfully connected to GitHub"
  }
}
```

---

### Sync Integration

Manually trigger sync.

**Endpoint:** `POST /api/v1/integrations/{integration_id}/sync`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Sync started",
  "data": {
    "sync_id": "sync-123"
  }
}
```

---

## Webhooks

### Create Webhook

**Endpoint:** `POST /api/v1/webhooks`

**Request Body:**
```json
{
  "organization_id": "org-123",
  "project_id": "project-123",
  "name": "CI/CD Webhook",
  "url": "https://api.example.com/webhook",
  "secret": "webhook-secret-key",
  "events": [
    "issue.created",
    "issue.updated",
    "issue.deleted",
    "comment.created"
  ],
  "ssl_verify": true,
  "timeout_seconds": 30
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "webhook": {
      "id": "webhook-123",
      "name": "CI/CD Webhook",
      "url": "https://api.example.com/webhook",
      "events": ["issue.created", "issue.updated"],
      "is_active": true,
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

### List Webhooks

**Endpoint:** `GET /api/v1/webhooks`

**Query Parameters:**
- `organization_id`: Required
- `project_id`: Optional

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "webhooks": [
      {
        "id": "webhook-123",
        "name": "CI/CD Webhook",
        "url": "https://api.example.com/webhook",
        "is_active": true,
        "last_triggered_at": "2026-03-06T09:45:00Z",
        "failure_count": 0
      }
    ],
    "total": 2
  }
}
```

---

### Update Webhook

**Endpoint:** `PATCH /api/v1/webhooks/{webhook_id}`

**Response:** `200 OK`

---

### Delete Webhook

**Endpoint:** `DELETE /api/v1/webhooks/{webhook_id}`

**Response:** `200 OK`

---

### Test Webhook

Send test payload to webhook.

**Endpoint:** `POST /api/v1/webhooks/{webhook_id}/test`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "delivered": true,
    "status_code": 200,
    "response_time_ms": 145
  }
}
```

---

### Webhook Deliveries

Get delivery logs for a webhook.

**Endpoint:** `GET /api/v1/webhooks/{webhook_id}/deliveries`

**Query Parameters:**
- `limit`: Results limit (default: 50)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "deliveries": [
      {
        "id": "delivery-123",
        "event_type": "issue.created",
        "delivered": true,
        "response_code": 200,
        "duration_ms": 145,
        "attempt_number": 1,
        "created_at": "2026-03-06T09:45:00Z"
      }
    ],
    "pagination": {
      "total": 234,
      "limit": 50,
      "offset": 0
    }
  }
}
```

---

### Redeliver Webhook

Retry failed webhook delivery.

**Endpoint:** `POST /api/v1/webhooks/deliveries/{delivery_id}/redeliver`

**Response:** `200 OK`

---

## Notifications

### Get Notifications

**Endpoint:** `GET /api/v1/notifications`

**Query Parameters:**
- `is_read`: Filter by read status (true/false)
- `type`: Filter by notification type
- `limit`: Results limit (default: 50)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "notifications": [
      {
        "id": "notif-123",
        "type": "issue_assigned",
        "title": "Issue assigned to you",
        "message": "John Doe assigned PROJ-123 to you",
        "link_url": "/issues/PROJ-123",
        "entity_type": "issue",
        "entity_id": "issue-123",
        "is_read": false,
        "priority": "normal",
        "created_at": "2026-03-06T10:00:00Z"
      }
    ],
    "pagination": {
      "total": 45,
      "unread_count": 12,
      "limit": 50,
      "offset": 0
    }
  }
}
```

---

### Mark Notification as Read

**Endpoint:** `PATCH /api/v1/notifications/{notification_id}/read`

**Response:** `200 OK`

---

### Mark All as Read

**Endpoint:** `POST /api/v1/notifications/mark-all-read`

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "All notifications marked as read",
  "data": {
    "marked_count": 12
  }
}
```

---

### Delete Notification

**Endpoint:** `DELETE /api/v1/notifications/{notification_id}`

**Response:** `200 OK`

---

### Get Notification Settings

**Endpoint:** `GET /api/v1/notifications/settings`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "settings": {
      "email_notifications": true,
      "push_notifications": true,
      "notification_types": {
        "issue_assigned": {
          "email": true,
          "push": true,
          "in_app": true
        },
        "comment_mentioned": {
          "email": true,
          "push": true,
          "in_app": true
        }
      }
    }
  }
}
```

---

### Update Notification Settings

**Endpoint:** `PATCH /api/v1/notifications/settings`

**Request Body:**
```json
{
  "email_notifications": false,
  "notification_types": {
    "issue_assigned": {
      "email": false,
      "push": true
    }
  }
}
```

**Response:** `200 OK`
