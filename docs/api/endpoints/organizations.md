# Organizations API Endpoints

**Base URL:** `/api/v1/organizations`  
**Version:** 1.0  
**Authentication:** Required

---

## Create Organization

Create a new organization.

**Endpoint:** `POST /api/v1/organizations`

**Request Body:**
```json
{
  "name": "Acme Corporation",
  "slug": "acme-corp",
  "description": "Leading software company",
  "website_url": "https://acme.com",
  "billing_email": "billing@acme.com"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "organization": {
      "id": "org-123",
      "name": "Acme Corporation",
      "slug": "acme-corp",
      "description": "Leading software company",
      "logo_url": null,
      "website_url": "https://acme.com",
      "plan_type": "free",
      "max_users": 10,
      "max_projects": 3,
      "storage_limit_bytes": 1073741824,
      "storage_used_bytes": 0,
      "created_at": "2026-03-06T10:00:00Z"
    }
  }
}
```

---

## List Organizations

Get organizations for authenticated user.

**Endpoint:** `GET /api/v1/organizations`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "organizations": [
      {
        "id": "org-123",
        "name": "Acme Corporation",
        "slug": "acme-corp",
        "logo_url": "https://cdn.mytodo.com/logos/acme.jpg",
        "role": "owner",
        "member_count": 15,
        "project_count": 8
      }
    ],
    "total": 1
  }
}
```

---

## Get Organization

Get organization details.

**Endpoint:** `GET /api/v1/organizations/{org_id}`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "organization": {
      "id": "org-123",
      "name": "Acme Corporation",
      "slug": "acme-corp",
      "description": "Leading software company",
      "logo_url": "https://cdn.mytodo.com/logos/acme.jpg",
      "website_url": "https://acme.com",
      "plan_type": "professional",
      "max_users": 100,
      "max_projects": 50,
      "storage_limit_bytes": 107374182400,
      "storage_used_bytes": 25769803776,
      "settings": {
        "require_2fa": true,
        "allowed_domains": ["acme.com"]
      },
      "created_at": "2026-01-01T00:00:00Z"
    }
  }
}
```

---

## Update Organization

Update organization details.

**Endpoint:** `PATCH /api/v1/organizations/{org_id}`

**Permissions:** Owner or Admin

**Request Body:**
```json
{
  "name": "Acme Corp",
  "description": "Updated description",
  "website_url": "https://acmecorp.com",
  "settings": {
    "require_2fa": true
  }
}
```

**Response:** `200 OK`

---

## Delete Organization

Delete organization (requires owner).

**Endpoint:** `DELETE /api/v1/organizations/{org_id}`

**Request Body:**
```json
{
  "confirmation": "DELETE"
}
```

**Response:** `200 OK`

---

## Organization Members

### List Members

**Endpoint:** `GET /api/v1/organizations/{org_id}/members`

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
        "role": "owner",
        "status": "active",
        "joined_at": "2026-01-01T00:00:00Z"
      }
    ],
    "total": 15
  }
}
```

---

### Invite Member

**Endpoint:** `POST /api/v1/organizations/{org_id}/members/invite`

**Request Body:**
```json
{
  "email": "newuser@example.com",
  "role": "member"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Invitation sent to newuser@example.com"
}
```

---

### Update Member Role

**Endpoint:** `PATCH /api/v1/organizations/{org_id}/members/{member_id}`

**Request Body:**
```json
{
  "role": "admin"
}
```

**Response:** `200 OK`

---

### Remove Member

**Endpoint:** `DELETE /api/v1/organizations/{org_id}/members/{member_id}`

**Response:** `200 OK`

---

## Organization Statistics

**Endpoint:** `GET /api/v1/organizations/{org_id}/stats`

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "stats": {
      "total_members": 15,
      "total_projects": 8,
      "total_issues": 234,
      "open_issues": 89,
      "storage_used_gb": 24,
      "storage_limit_gb": 100,
      "active_users_this_month": 12
    }
  }
}
```
