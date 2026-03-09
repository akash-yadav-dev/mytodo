import { API_BASE_URL } from "@/lib/constants/api";
import { getAccessToken } from "@/services/auth/token";

type RequestOptions = RequestInit & { shouldAuth?: boolean };

export async function apiFetch<T>(
  path: string,
  options: RequestOptions = {}
): Promise<T> {
  const { shouldAuth = true, headers, ...rest } = options;
  const token = shouldAuth ? getAccessToken() : null;

  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...rest,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(headers ?? {}),
    },
  });

  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || "Request failed");
  }

  if (response.status === 204) {
    return {} as T;
  }

  return (await response.json()) as T;
}
