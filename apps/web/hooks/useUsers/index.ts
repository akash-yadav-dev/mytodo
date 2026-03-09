"use client";

import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  getCurrentUserProfile,
  getUserProfileById,
  listUserProfiles,
  searchUserProfiles,
  createUserProfile,
  updateUserProfile,
  deleteUserProfile,
  getUserPreferences,
  updateUserPreferences,
} from "@/services/api/users";
import type {
  CreateUserProfilePayload,
  UpdateUserProfilePayload,
  UpdateUserPreferencesPayload,
  UserProfile,
  UserPreferences,
} from "@/lib/types/users";

// Query keys for cache management
export const userKeys = {
  all: ["users"] as const,
  lists: () => [...userKeys.all, "list"] as const,
  list: (page: number, limit: number) => [...userKeys.lists(), { page, limit }] as const,
  details: () => [...userKeys.all, "detail"] as const,
  detail: (id: string) => [...userKeys.details(), id] as const,
  current: () => [...userKeys.all, "current"] as const,
  preferences: () => [...userKeys.all, "preferences"] as const,
  search: (query: string) => [...userKeys.all, "search", query] as const,
};

/**
 * Hook to get the current user's profile
 */
export function useCurrentUserProfile() {
  return useQuery({
    queryKey: userKeys.current(),
    queryFn: getCurrentUserProfile,
    staleTime: 5 * 60 * 1000, // 5 minutes
  });
}

/**
 * Hook to get a user profile by ID
 */
export function useUserProfile(userId: string) {
  return useQuery({
    queryKey: userKeys.detail(userId),
    queryFn: () => getUserProfileById(userId),
    enabled: !!userId,
    staleTime: 5 * 60 * 1000,
  });
}

/**
 * Hook to list user profiles with pagination
 */
export function useUserProfiles(page: number = 1, limit: number = 10) {
  return useQuery({
    queryKey: userKeys.list(page, limit),
    queryFn: () => listUserProfiles(page, limit),
    staleTime: 2 * 60 * 1000, // 2 minutes
  });
}

/**
 * Hook to search user profiles
 */
export function useUserSearch(query: string, limit: number = 20) {
  return useQuery({
    queryKey: userKeys.search(query),
    queryFn: () => searchUserProfiles(query, limit),
    enabled: query.length >= 2,
    staleTime: 1 * 60 * 1000, // 1 minute
  });
}

/**
 * Hook to get user preferences
 */
export function useUserPreferences() {
  return useQuery({
    queryKey: userKeys.preferences(),
    queryFn: getUserPreferences,
    staleTime: 5 * 60 * 1000,
  });
}

/**
 * Hook to create a user profile
 */
export function useCreateUserProfile() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateUserProfilePayload) => createUserProfile(payload),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: userKeys.current() });
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
    },
  });
}

/**
 * Hook to update user profile
 */
export function useUpdateUserProfile() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: UpdateUserProfilePayload) => updateUserProfile(payload),
    onSuccess: (data: UserProfile) => {
      queryClient.setQueryData(userKeys.current(), data);
      queryClient.invalidateQueries({ queryKey: userKeys.detail(data.user_id) });
      queryClient.invalidateQueries({ queryKey: userKeys.lists() });
    },
  });
}

/**
 * Hook to delete user profile
 */
export function useDeleteUserProfile() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: deleteUserProfile,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: userKeys.all });
    },
  });
}

/**
 * Hook to update user preferences
 */
export function useUpdateUserPreferences() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: UpdateUserPreferencesPayload) => updateUserPreferences(payload),
    onSuccess: (data: UserPreferences) => {
      queryClient.setQueryData(userKeys.preferences(), data);
    },
  });
}
