import type { 
  UserProfile, 
  UserPreferences, 
  UpdateUserProfilePayload, 
  UpdateUserPreferencesPayload,
  CreateUserProfilePayload,
  PaginatedUserProfiles
} from "@/lib/types/users";
import { apiFetch } from "@/services/api/client";

/**
 * Get the current authenticated user's profile
 */
export async function getCurrentUserProfile(): Promise<UserProfile> {
  const response = await apiFetch<{ success: boolean; data: UserProfile }>(
    "/api/v1/users/me",
    { method: "GET" }
  );
  return response.data;
}

/**
 * Get a user profile by ID
 */
export async function getUserProfileById(userId: string): Promise<UserProfile> {
  const response = await apiFetch<{ success: boolean; data: UserProfile }>(
    `/api/v1/users/${userId}`,
    { method: "GET", shouldAuth: false }
  );
  return response.data;
}

/**
 * List all user profiles with pagination
 */
export async function listUserProfiles(page: number = 1, limit: number = 10): Promise<PaginatedUserProfiles> {
  const response = await apiFetch<{ success: boolean; data: PaginatedUserProfiles }>(
    `/api/v1/users?page=${page}&limit=${limit}`,
    { method: "GET", shouldAuth: false }
  );
  return response.data;
}

/**
 * Search user profiles
 */
export async function searchUserProfiles(query: string, limit: number = 20): Promise<UserProfile[]> {
  const response = await apiFetch<{ success: boolean; data: UserProfile[] }>(
    `/api/v1/users/search?q=${encodeURIComponent(query)}&limit=${limit}`,
    { method: "GET", shouldAuth: false }
  );
  return response.data;
}

/**
 * Create a user profile (typically called after registration)
 */
export async function createUserProfile(payload: CreateUserProfilePayload): Promise<UserProfile> {
  const response = await apiFetch<{ success: boolean; data: UserProfile }>(
    "/api/v1/users/profile",
    {
      method: "POST",
      body: JSON.stringify(payload),
    }
  );
  return response.data;
}

/**
 * Update the current user's profile
 */
export async function updateUserProfile(payload: UpdateUserProfilePayload): Promise<UserProfile> {
  const response = await apiFetch<{ success: boolean; data: UserProfile }>(
    "/api/v1/users/me",
    {
      method: "PUT",
      body: JSON.stringify(payload),
    }
  );
  return response.data;
}

/**
 * Delete the current user's profile
 */
export async function deleteUserProfile(): Promise<void> {
  await apiFetch("/api/v1/users/me", { method: "DELETE" });
}

/**
 * Get the current user's preferences
 */
export async function getUserPreferences(): Promise<UserPreferences> {
  const response = await apiFetch<{ success: boolean; data: UserPreferences }>(
    "/api/v1/users/me/preferences",
    { method: "GET" }
  );
  return response.data;
}

/**
 * Update the current user's preferences
 */
export async function updateUserPreferences(payload: UpdateUserPreferencesPayload): Promise<UserPreferences> {
  const response = await apiFetch<{ success: boolean; data: UserPreferences }>(
    "/api/v1/users/me/preferences",
    {
      method: "PUT",
      body: JSON.stringify(payload),
    }
  );
  return response.data;
}
