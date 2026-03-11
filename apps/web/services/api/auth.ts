import { apiFetch } from "@/services/api/client";
import type { AuthApiResponse, AuthSession, AuthUser, LoginPayload, RegisterPayload } from "@/lib/types/auth";

function normalizeAuthResponse(raw: AuthApiResponse): AuthSession {
  return {
    accessToken: raw.access_token,
    refreshToken: raw.refresh_token,
    user: raw.user,
  };
}

export async function loginUser(payload: LoginPayload): Promise<AuthSession> {
  const raw = await apiFetch<AuthApiResponse>("/api/v1/auth/login", {
    method: "POST",
    body: JSON.stringify(payload),
    shouldAuth: false,
  });
  return normalizeAuthResponse(raw);
}

export async function registerUser(payload: RegisterPayload): Promise<AuthSession> {
  const raw = await apiFetch<AuthApiResponse>("/api/v1/auth/register", {
    method: "POST",
    body: JSON.stringify(payload),
    shouldAuth: false,
  });
  return normalizeAuthResponse(raw);
}

export async function fetchCurrentUser(): Promise<AuthUser | null> {
  try {
    return await apiFetch<AuthUser>("/api/v1/auth/me");
  } catch {
    return null;
  }
}

export async function logoutUser(): Promise<void> {
  await apiFetch("/api/v1/auth/logout", { method: "POST" });
}

export async function refreshAccessToken(refreshToken: string): Promise<AuthSession> {
  const raw = await apiFetch<AuthApiResponse>("/api/v1/auth/refresh", {
    method: "POST",
    body: JSON.stringify({ refresh_token: refreshToken }),
    shouldAuth: false,
  });
  return normalizeAuthResponse(raw);
}
