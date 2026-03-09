import type { AuthSession, AuthUser, LoginPayload, RegisterPayload } from "@/lib/types/auth";
import { apiFetch } from "@/services/api/client";

export async function loginUser(payload: LoginPayload): Promise<AuthSession> {
  return apiFetch<AuthSession>("/api/v1/auth/login", {
    method: "POST",
    body: JSON.stringify(payload),
    shouldAuth: false,
  });
}

export async function registerUser(
  payload: RegisterPayload
): Promise<AuthSession> {
  return apiFetch<AuthSession>("/api/v1/auth/register", {
    method: "POST",
    body: JSON.stringify(payload),
    shouldAuth: false,
  });
}

export async function fetchCurrentUser(): Promise<AuthUser | null> {
  try {
    return await apiFetch<AuthUser>("/api/v1/auth/me", {
      method: "GET",
    });
  } catch {
    return null;
  }
}

export async function logoutUser(): Promise<void> {
  await apiFetch("/api/v1/auth/logout", { method: "POST" });
}
