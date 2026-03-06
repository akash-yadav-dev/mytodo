# Comments & Attachments API Endpoints

**Base URL:** `/api/v1`  
**Version:** 1.0  
**Authentication:** Required

---

## Comments

### Create Comment

Add comment to an issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/comments`

**Request Body:**
```json
{
  "content": "This looks good to me. Suggested using async/await instead of promises.",
  "is_internal": false
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "comment": {
      "id": "comment-123",
      "issue_id": "issue-123",
      "user": {
        "id": "user-123",
        "username": "johndoe",
        "full_name": "John Doe",
        "avatar_url": "url"
      },
      "content": "This looks good to me...",
      "content_html": "<p>This looks good to me...</p>",
      "is_internal": false,
      "edited": false,
      "created_at": "2026-03-06T10:00:00Z",
      "updated_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

### List Comments

Get all comments for an issue.

**Endpoint:** `GET /api/v1/issues/{issue_id}/comments`

**Query Parameters:**
- `limit`: Results limit (default: 50)
- `offset`: Pagination offset

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "comments": [
      {
        "id": "comment-123",
        "user": {
          "id": "user-123",
          "username": "johndoe",
          "full_name": "John Doe",
          "avatar_url": "url"
        },
        "content": "Comment text",
        "is_internal": false,
        "edited": false,
        "created_at": "2026-03-06T10:00:00Z"
      }
    ],
    "total": 5
  }
}
```

---

### Update Comment

Edit an existing comment.

**Endpoint:** `PATCH /api/v1/comments/{comment_id}`

**Permissions:** Comment author or admin

**Request Body:**
```json
{
  "content": "Updated comment text"
}
```

**Response:** `200 OK`

---

### Delete Comment

Delete a comment.

**Endpoint:** `DELETE /api/v1/comments/{comment_id}`

**Permissions:** Comment author or admin

**Response:** `200 OK`

---

### Reply to Comment

Create a threaded reply.

**Endpoint:** `POST /api/v1/comments/{comment_id}/replies`

**Request Body:**
```json
{
  "content": "Reply text"
}
```

**Response:** `201 Created`

---

## Attachments

### Upload Attachment

Upload file to an issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/attachments`

**Request:** `multipart/form-data`
```
file: [File]
```

**Max File Size:** 50MB per file

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "attachment": {
      "id": "attachment-123",
      "issue_id": "issue-123",
      "file_name": "screenshot.png",
      "file_size": 2048576,
      "file_type": "image/png",
      "storage_key": "attachments/2026/03/06/uuid-screenshot.png",
      "thumbnail_url": "https://cdn.mytodo.com/thumbnails/uuid.jpg",
      "download_url": "https://cdn.mytodo.com/attachments/uuid-screenshot.png",
      "uploaded_by": {
        "id": "user-123",
        "full_name": "John Doe"
      },
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

### List Attachments

Get all attachments for an issue.

**Endpoint:** `GET /api/v1/issues/{issue_id}/attachments`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "attachments": [
      {
        "id": "attachment-123",
        "file_name": "screenshot.png",
        "file_size": 2048576,
        "file_type": "image/png",
        "thumbnail_url": "https://cdn.mytodo.com/thumbnails/uuid.jpg",
        "download_url": "https://cdn.mytodo.com/attachments/uuid.png",
        "download_count": 5,
        "uploaded_by": {
          "id": "user-123",
          "full_name": "John Doe",
          "avatar_url": "url"
        },
        "created_at": "2026-03-06T10:00:00Z"
      }
    ],
    "total": 3,
    "total_size_bytes": 5242880
  }
}
```

---

### Download Attachment

Download attachment file.

**Endpoint:** `GET /api/v1/attachments/{attachment_id}/download`

**Response:** `200 OK`
- Content-Type: `application/octet-stream` or specific file type
- Content-Disposition: `attachment; filename="screenshot.png"`

---

### Delete Attachment

Delete an attachment.

**Endpoint:** `DELETE /api/v1/attachments/{attachment_id}`

**Permissions:** Uploader or admin

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Attachment deleted successfully"
}
```

---

## Worklogs

### Log Time

Log time spent on an issue.

**Endpoint:** `POST /api/v1/issues/{issue_id}/worklogs`

**Request Body:**
```json
{
  "time_spent_minutes": 120,
  "comment": "Implemented user authentication",
  "started_at": "2026-03-06T09:00:00Z"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "worklog": {
      "id": "worklog-123",
      "issue_id": "issue-123",
      "user": {
        "id": "user-123",
        "full_name": "John Doe"
      },
      "time_spent_minutes": 120,
      "comment": "Implemented user authentication",
      "started_at": "2026-03-06T09:00:00Z",
      "created_at": "2026-03-06T11:00:00Z"
    }
  }
}
```

---

### List Worklogs

Get worklogs for an issue.

**Endpoint:** `GET /api/v1/issues/{issue_id}/worklogs`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "worklogs": [
      {
        "id": "worklog-123",
        "user": {
          "id": "user-123",
          "full_name": "John Doe",
          "avatar_url": "url"
        },
        "time_spent_minutes": 120,
        "comment": "Implemented user authentication",
        "started_at": "2026-03-06T09:00:00Z",
        "created_at": "2026-03-06T11:00:00Z"
      }
    ],
    "total": 8,
    "total_time_minutes": 960
  }
}
```

---

### Update Worklog

**Endpoint:** `PATCH /api/v1/worklogs/{worklog_id}`

**Request Body:**
```json
{
  "time_spent_minutes": 150,
  "comment": "Updated description"
}
```

**Response:** `200 OK`

---

### Delete Worklog

**Endpoint:** `DELETE /api/v1/worklogs/{worklog_id}`

**Response:** `200 OK`
