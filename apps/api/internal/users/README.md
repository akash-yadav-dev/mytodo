# Users Module Documentation

## Overview

The Users module manages user profiles and preferences, extending the authentication module with rich user information and customization options.

## Architecture

This module follows Clean Architecture principles with clear separation of concerns:

```
users/
├── application/          # Application layer (use cases)
│   ├── commands/        # Write operations (CQRS)
│   ├── handlers/        # Command and query handlers
│   └── queries/         # Read operations (CQRS)
├── domain/              # Domain layer (business logic)
│   ├── entity/          # Domain entities
│   ├── repository/      # Repository interfaces
│   └── service/         # Domain services
├── infrastructure/      # Infrastructure layer
│   ├── cache/          # Caching implementations
│   └── persistence/    # Database implementations
└── interfaces/          # Interface adapters
    ├── dto/            # Data Transfer Objects
    └── http/           # HTTP controllers and routes
```

## Database Schema

### Tables

#### user_profiles
Stores extended user profile information.

```sql
CREATE TABLE user_profiles (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE REFERENCES users(id),
    username VARCHAR(50) UNIQUE,
    display_name VARCHAR(100),
    avatar_url TEXT,
    bio TEXT,
    location VARCHAR(100),
    website VARCHAR(255),
    phone VARCHAR(20),
    timezone VARCHAR(50) DEFAULT 'UTC',
    language VARCHAR(10) DEFAULT 'en',
    theme VARCHAR(20) DEFAULT 'light',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

#### user_preferences
Stores user notification and application preferences.

```sql
CREATE TABLE user_preferences (
    id UUID PRIMARY KEY,
    user_id UUID UNIQUE REFERENCES users(id),
    email_notifications BOOLEAN DEFAULT true,
    push_notifications BOOLEAN DEFAULT true,
    newsletter_subscription BOOLEAN DEFAULT false,
    weekly_digest BOOLEAN DEFAULT true,
    mentions_notifications BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

## API Endpoints

### Public Endpoints

#### List Users
```http
GET /api/v1/users?page=1&limit=10
```

Response:
```json
{
  "success": true,
  "data": {
    "users": [...],
    "total": 100,
    "page": 1,
    "limit": 10
  }
}
```

#### Get User Profile by ID
```http
GET /api/v1/users/:id
```

Response:
```json
{
  "success": true,
  "data": {
    "id": "uuid",
    "user_id": "uuid",
    "username": "johndoe",
    "display_name": "John Doe",
    "avatar_url": "https://...",
    "bio": "Software engineer",
    "location": "New York, NY",
    "timezone": "America/New_York",
    "language": "en",
    "theme": "dark"
  }
}
```

#### Search Users
```http
GET /api/v1/users/search?q=john&limit=20
```

### Protected Endpoints (Require Authentication)

#### Get Current User Profile
```http
GET /api/v1/users/me
Authorization: Bearer <token>
```

#### Create User Profile
```http
POST /api/v1/users/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "username": "johndoe",
  "display_name": "John Doe",
  "avatar_url": "https://..."
}
```

#### Update User Profile
```http
PUT /api/v1/users/me
Authorization: Bearer <token>
Content-Type: application/json

{
  "display_name": "John Doe",
  "bio": "Updated bio",
  "location": "San Francisco, CA",
  "website": "https://johndoe.com",
  "theme": "dark"
}
```

#### Delete User Profile
```http
DELETE /api/v1/users/me
Authorization: Bearer <token>
```

#### Get User Preferences
```http
GET /api/v1/users/me/preferences
Authorization: Bearer <token>
```

#### Update User Preferences
```http
PUT /api/v1/users/me/preferences
Authorization: Bearer <token>
Content-Type: application/json

{
  "email_notifications": true,
  "push_notifications": false,
  "newsletter_subscription": true,
  "weekly_digest": true,
  "mentions_notifications": true
}
```

## Frontend Integration

### React Hooks

```typescript
import {
  useCurrentUserProfile,
  useUserProfile,
  useUserProfiles,
  useUpdateUserProfile,
  useUserPreferences,
  useUpdateUserPreferences
} from '@/hooks/useUsers';

// Get current user profile
const { data: profile, isLoading } = useCurrentUserProfile();

// Get specific user profile
const { data: userProfile } = useUserProfile(userId);

// List users with pagination
const { data: users } = useUserProfiles(page, limit);

// Update profile
const updateProfile = useUpdateUserProfile();
await updateProfile.mutateAsync({
  display_name: "New Name",
  bio: "Updated bio"
});

// Get and update preferences
const { data: prefs } = useUserPreferences();
const updatePrefs = useUpdateUserPreferences();
await updatePrefs.mutateAsync({
  email_notifications: false
});
```

### TypeScript Types

```typescript
import type {
  UserProfile,
  UserPreferences,
  UpdateUserProfilePayload,
  UpdateUserPreferencesPayload
} from '@/lib/types/users';
```

## Running Migrations

### Apply Migrations
```bash
./infrastructure/scripts/run_migrations.sh up
```

### Rollback Migration
```bash
./infrastructure/scripts/run_migrations.sh down 1
```

### Create New Migration
```bash
./infrastructure/scripts/run_migrations.sh create add_user_field
```

### Check Migration Version
```bash
./infrastructure/scripts/run_migrations.sh version
```

## Seeding Data

Run the seed script to populate sample user profiles:

```bash
psql -U postgres -d mytodo -f infrastructure/scripts/seed_users.sql
```

## Integration with Auth Module

The Users module integrates seamlessly with the Auth module:

1. **Registration Flow**: When a user registers via `/api/v1/auth/register`, the auth service creates the user account, and a user profile is automatically created
2. **Profile Creation**: After authentication, users can create/update their profile via `/api/v1/users/profile`
3. **User Context**: The JWT token contains the user ID, which is used to fetch the corresponding profile

### UserRegistrationService

The `UserRegistrationService` coordinates between auth and users modules:

```go
// Register user with profile creation
user, tokens, err := registrationService.RegisterUserWithProfile(
    ctx,
    email,
    password,
    name,
)
```

## Development

### Running Tests

```bash
# Unit tests
go test ./apps/api/internal/users/...

# Integration tests
go test -tags=integration ./apps/api/internal/users/...
```

### Code Generation

If you need to regenerate DTOs or mocks:

```bash
go generate ./apps/api/internal/users/...
```

## Best Practices

1. **Always validate input** in command/query objects
2. **Use DTOs** for API responses, never expose domain entities directly
3. **Handle errors gracefully** with appropriate HTTP status codes
4. **Use transactions** for operations that modify multiple tables
5. **Cache frequently accessed data** like user profiles
6. **Log important events** like profile updates, deletions
7. **Sanitize user input** especially for bio, website fields
8. **Rate limit** profile updates to prevent abuse

## Troubleshooting

### Profile not found after registration
- Check if user profile creation succeeded in logs
- Verify user_id foreign key constraint
- Ensure migrations ran successfully

### Username already taken error
- Username must be unique across all profiles
- Check for case-sensitivity issues
- Use search to verify username availability

### Preferences not updating
- Verify JWT token is valid and contains correct user_id
- Check database constraints
- Ensure preferences record exists (created automatically with profile)

## Future Enhancements

- [ ] Add profile picture upload to cloud storage
- [ ] Implement profile visibility settings (public/private)
- [ ] Add social media link validation
- [ ] Implement user blocking/following
- [ ] Add profile completion percentage
- [ ] Implement profile badges/achievements
- [ ] Add activity history tracking
