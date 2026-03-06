# Authentication API Endpoints

**Base URL:** `/api/v1/auth`  
**Version:** 1.0  
**Last Updated:** March 6, 2026

This document describes all authentication-related API endpoints.

---

## Table of Contents

1. [Registration](#registration)
2. [Login](#login)
3. [OAuth](#oauth)
4. [Token Management](#token-management)
5. [Password Management](#password-management)
6. [Two-Factor Authentication](#two-factor-authentication)
7. [Session Management](#session-management)

---

## Registration

### Register New User

Create a new user account.

**Endpoint:** `POST /api/v1/auth/register`

**Request Body:**
```json
{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "SecurePass123!",
  "full_name": "John Doe",
  "timezone": "America/New_York",
  "locale": "en-US"
}
```

**Response:** `201 Created`
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@example.com",
      "username": "johndoe",
      "full_name": "John Doe",
      "avatar_url": null,
      "email_verified": false,
      "created_at": "2026-03-06T10:00:00Z"
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIs...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
      "token_type": "Bearer",
      "expires_in": 3600
    }
  },
  "message": "Registration successful. Please verify your email."
}
```

**Validation Rules:**
- Email: Valid email format, unique
- Username: 3-50 characters, alphanumeric with underscores, unique
- Password: Minimum 8 characters, at least one uppercase, one lowercase, one number
- Full name: 1-255 characters

**Error Responses:**

`409 Conflict` - Email or username already exists
```json
{
  "success": false,
  "error": {
    "code": "DUPLICATE_EMAIL",
    "message": "Email address already registered",
    "field": "email"
  }
}
```

`400 Bad Request` - Validation errors
```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Validation failed",
    "details": [
      {
        "field": "password",
        "message": "Password must be at least 8 characters"
      }
    ]
  }
}
```

---

### Verify Email

Verify user email address.

**Endpoint:** `POST /api/v1/auth/verify-email`

**Request Body:**
```json
{
  "token": "verification-token-from-email"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Email verified successfully"
}
```

---

### Resend Verification Email

Request a new verification email.

**Endpoint:** `POST /api/v1/auth/resend-verification`

**Request Body:**
```json
{
  "email": "user@example.com"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Verification email sent"
}
```

---

## Login

### Login with Email/Password

Authenticate user with credentials.

**Endpoint:** `POST /api/v1/auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePass123!",
  "remember_me": true
}
```

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
      "email_verified": true,
      "two_factor_enabled": false,
      "preferences": {
        "theme": "light",
        "notifications": true
      }
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIs...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
      "token_type": "Bearer",
      "expires_in": 3600
    },
    "session": {
      "id": "session-id-123",
      "expires_at": "2026-03-07T10:00:00Z"
    }
  }
}
```

**Error Responses:**

`401 Unauthorized` - Invalid credentials
```json
{
  "success": false,
  "error": {
    "code": "INVALID_CREDENTIALS",
    "message": "Invalid email or password"
  }
}
```

`423 Locked` - Account locked
```json
{
  "success": false,
  "error": {
    "code": "ACCOUNT_LOCKED",
    "message": "Account temporarily locked due to failed login attempts",
    "locked_until": "2026-03-06T11:00:00Z"
  }
}
```

`403 Forbidden` - Account suspended
```json
{
  "success": false,
  "error": {
    "code": "ACCOUNT_SUSPENDED",
    "message": "Your account has been suspended. Contact support."
  }
}
```

---

### Login with Username

Alternative login with username.

**Endpoint:** `POST /api/v1/auth/login/username`

**Request Body:**
```json
{
  "username": "johndoe",
  "password": "SecurePass123!"
}
```

**Response:** Same as email login

---

## OAuth

### Get OAuth URL

Get authorization URL for OAuth provider.

**Endpoint:** `GET /api/v1/auth/oauth/{provider}/url`

**Path Parameters:**
- `provider`: OAuth provider (google, github, microsoft, gitlab)

**Query Parameters:**
- `redirect_uri`: Callback URL after authentication
- `state`: Optional state parameter for CSRF protection

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "authorization_url": "https://accounts.google.com/o/oauth2/v2/auth?client_id=...",
    "state": "random-state-token"
  }
}
```

---

### OAuth Callback

Handle OAuth provider callback.

**Endpoint:** `GET /api/v1/auth/oauth/{provider}/callback`

**Query Parameters:**
- `code`: Authorization code from provider
- `state`: State parameter for verification

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "user": {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "email": "user@gmail.com",
      "username": "johndoe",
      "full_name": "John Doe",
      "avatar_url": "https://lh3.googleusercontent.com/...",
      "email_verified": true
    },
    "tokens": {
      "access_token": "eyJhbGciOiJIUzI1NiIs...",
      "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
      "token_type": "Bearer",
      "expires_in": 3600
    },
    "is_new_user": false
  }
}
```

---

### Link OAuth Provider

Link OAuth provider to existing account.

**Endpoint:** `POST /api/v1/auth/oauth/{provider}/link`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "code": "authorization-code-from-provider"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Google account linked successfully"
}
```

---

### Unlink OAuth Provider

Remove OAuth provider connection.

**Endpoint:** `DELETE /api/v1/auth/oauth/{provider}`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Google account unlinked successfully"
}
```

---

## Token Management

### Refresh Token

Get new access token using refresh token.

**Endpoint:** `POST /api/v1/auth/token/refresh`

**Request Body:**
```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIs...",
    "token_type": "Bearer",
    "expires_in": 3600
  }
}
```

---

### Validate Token

Validate access token.

**Endpoint:** `POST /api/v1/auth/token/validate`

**Request Body:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "valid": true,
    "expires_at": "2026-03-06T11:00:00Z",
    "user_id": "550e8400-e29b-41d4-a716-446655440000"
  }
}
```

---

### Logout

Invalidate current session.

**Endpoint:** `POST /api/v1/auth/logout`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Logged out successfully"
}
```

---

### Logout All Devices

Invalidate all user sessions.

**Endpoint:** `POST /api/v1/auth/logout/all`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "All sessions terminated",
  "sessions_revoked": 3
}
```

---

## Password Management

### Request Password Reset

Request password reset email.

**Endpoint:** `POST /api/v1/auth/password/reset-request`

**Request Body:**
```json
{
  "email": "user@example.com"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Password reset instructions sent to your email"
}
```

**Note:** Always returns 200 even if email doesn't exist (security)

---

### Verify Reset Token

Verify password reset token validity.

**Endpoint:** `POST /api/v1/auth/password/verify-token`

**Request Body:**
```json
{
  "token": "reset-token-from-email"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "valid": true,
    "email": "u***@example.com",
    "expires_at": "2026-03-06T12:00:00Z"
  }
}
```

---

### Reset Password

Reset password with token.

**Endpoint:** `POST /api/v1/auth/password/reset`

**Request Body:**
```json
{
  "token": "reset-token-from-email",
  "password": "NewSecurePass123!",
  "password_confirmation": "NewSecurePass123!"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Password reset successfully"
}
```

---

### Change Password

Change password for authenticated user.

**Endpoint:** `POST /api/v1/auth/password/change`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "current_password": "OldPassword123!",
  "new_password": "NewPassword123!",
  "new_password_confirmation": "NewPassword123!"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Password changed successfully"
}
```

---

## Two-Factor Authentication

### Enable 2FA

Enable two-factor authentication.

**Endpoint:** `POST /api/v1/auth/2fa/enable`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "secret": "JBSWY3DPEHPK3PXP",
    "qr_code": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
    "backup_codes": [
      "12345-67890",
      "23456-78901",
      "34567-89012",
      "45678-90123",
      "56789-01234"
    ]
  },
  "message": "Scan QR code with authenticator app"
}
```

---

### Verify 2FA Setup

Verify and activate 2FA setup.

**Endpoint:** `POST /api/v1/auth/2fa/verify`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "code": "123456"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Two-factor authentication enabled successfully"
}
```

---

### Disable 2FA

Disable two-factor authentication.

**Endpoint:** `POST /api/v1/auth/2fa/disable`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "password": "CurrentPassword123!",
  "code": "123456"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Two-factor authentication disabled"
}
```

---

### Verify 2FA Code

Verify 2FA code during login.

**Endpoint:** `POST /api/v1/auth/2fa/verify-login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "Password123!",
  "code": "123456"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "user": { /* user object */ },
    "tokens": { /* tokens object */ }
  }
}
```

---

### Regenerate Backup Codes

Generate new backup codes.

**Endpoint:** `POST /api/v1/auth/2fa/backup-codes/regenerate`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Request Body:**
```json
{
  "password": "CurrentPassword123!"
}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "backup_codes": [
      "12345-67890",
      "23456-78901",
      "34567-89012",
      "45678-90123",
      "56789-01234"
    ]
  },
  "message": "Backup codes regenerated. Store them securely."
}
```

---

## Session Management

### Get Active Sessions

List all active user sessions.

**Endpoint:** `GET /api/v1/auth/sessions`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "data": {
    "sessions": [
      {
        "id": "session-1",
        "device_info": {
          "browser": "Chrome 120",
          "os": "Windows 11",
          "device_type": "desktop"
        },
        "ip_address": "192.168.1.1",
        "location": "New York, US",
        "last_activity": "2026-03-06T10:00:00Z",
        "created_at": "2026-03-05T08:00:00Z",
        "is_current": true
      },
      {
        "id": "session-2",
        "device_info": {
          "browser": "Safari",
          "os": "iOS 17",
          "device_type": "mobile"
        },
        "ip_address": "192.168.1.2",
        "location": "New York, US",
        "last_activity": "2026-03-05T22:00:00Z",
        "created_at": "2026-03-04T15:00:00Z",
        "is_current": false
      }
    ],
    "total": 2
  }
}
```

---

### Revoke Session

Terminate a specific session.

**Endpoint:** `DELETE /api/v1/auth/sessions/{session_id}`

**Headers:**
```
Authorization: Bearer {access_token}
```

**Response:** `200 OK`
```json
{
  "success": true,
  "message": "Session terminated successfully"
}
```

---

## Rate Limiting

All authentication endpoints have rate limiting:

| Endpoint | Limit |
|----------|-------|
| Registration | 5 requests per hour per IP |
| Login | 10 requests per 15 minutes per IP |
| Password Reset | 3 requests per hour per email |
| 2FA Verification | 5 requests per 5 minutes per user |
| Token Refresh | 60 requests per minute per user |

---

## Error Codes

| Code | HTTP Status | Description |
|------|-------------|-------------|
| `VALIDATION_ERROR` | 400 | Request validation failed |
| `INVALID_CREDENTIALS` | 401 | Invalid email/password |
| `UNAUTHORIZED` | 401 | Missing or invalid token |
| `FORBIDDEN` | 403 | Account suspended/deleted |
| `NOT_FOUND` | 404 | Resource not found |
| `DUPLICATE_EMAIL` | 409 | Email already exists |
| `DUPLICATE_USERNAME` | 409 | Username already exists |
| `ACCOUNT_LOCKED` | 423 | Account temporarily locked |
| `RATE_LIMIT_EXCEEDED` | 429 | Too many requests |
| `INTERNAL_ERROR` | 500 | Server error |

---

## Security Headers

All responses include:
```
Content-Security-Policy: default-src 'self'
X-Content-Type-Options: nosniff
X-Frame-Options: DENY
X-XSS-Protection: 1; mode=block
Strict-Transport-Security: max-age=31536000; includeSubDomains
```

---

## See Also

- [Users API](users.md)
- [Organizations API](organizations.md)
- [Security Best Practices](../security.md)
