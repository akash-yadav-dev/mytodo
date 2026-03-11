import { apiFetch } from "@/services/api/client";
import type { AuthUser } from "@/lib/types/auth";

export type UpdateUserPayload = {
  name?: string;
};

export async function getCurrentUser(): Promise<AuthUser> {
  return apiFetch<AuthUser>("/api/v1/auth/me");
}

export async function updateUser(payload: UpdateUserPayload): Promise<AuthUser> {
  return apiFetch<AuthUser>("/api/v1/users/me", {
    method: "PUT",
    body: JSON.stringify(payload),
  });
}
