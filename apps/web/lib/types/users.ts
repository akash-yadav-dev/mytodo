// User Profile Types

export interface UserProfile {
  id: string;
  user_id: string;
  username: string | null;
  display_name: string;
  avatar_url?: string;
  bio?: string;
  location?: string;
  website?: string;
  phone?: string;
  timezone: string;
  language: string;
  theme: string;
  created_at: string;
  updated_at: string;
}

export interface UserPreferences {
  id: string;
  user_id: string;
  email_notifications: boolean;
  push_notifications: boolean;
  newsletter_subscription: boolean;
  weekly_digest: boolean;
  mentions_notifications: boolean;
  created_at: string;
  updated_at: string;
}

export interface PaginatedUserProfiles {
  users: UserProfile[];
  total: number;
  page: number;
  limit: number;
}

// Request Payloads

export interface CreateUserProfilePayload {
  username?: string;
  display_name: string;
  avatar_url?: string;
}

export interface UpdateUserProfilePayload {
  username?: string;
  display_name?: string;
  bio?: string;
  location?: string;
  website?: string;
  avatar_url?: string;
  phone?: string;
  timezone?: string;
  language?: string;
  theme?: 'light' | 'dark';
}

export interface UpdateUserPreferencesPayload {
  email_notifications: boolean;
  push_notifications: boolean;
  newsletter_subscription: boolean;
  weekly_digest: boolean;
  mentions_notifications: boolean;
}

// Search and Filter Types

export interface UserSearchParams {
  query: string;
  limit?: number;
}

export interface UserListParams {
  page?: number;
  limit?: number;
}
