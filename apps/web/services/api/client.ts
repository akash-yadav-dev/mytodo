import { API_BASE_URL } from "@/lib/constants/api";
import { getAccessToken, setAccessToken, clearAccessToken, getRefreshToken, clearRefreshToken } from "@/services/auth/token";
import { clearStoredSession } from "@/services/auth/session";

// ─── Typed API error ─────────────────────────────────────────────────────────

export class ApiError extends Error {
  constructor(
    public readonly status: number,
    message: string
  ) {
    super(message);
    this.name = "ApiError";
  }
}

// ─── Token refresh state (singleton to avoid parallel refresh races) ─────────

let refreshPromise: Promise<string | null> | null = null;

async function doRefresh(): Promise<string | null> {
  const refresh = getRefreshToken();
  if (!refresh) return null;

  try {
    const res = await fetch(`${API_BASE_URL}/api/v1/auth/refresh`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ refresh_token: refresh }),
    });
    if (!res.ok) throw new Error("refresh failed");
    const data = (await res.json()) as { access_token: string };
    setAccessToken(data.access_token);
    return data.access_token;
  } catch {
    clearAccessToken();
    clearRefreshToken();
    clearStoredSession();
    return null;
  }
}

// ─── Core fetch wrapper ───────────────────────────────────────────────────────

type RequestOptions = RequestInit & {
  shouldAuth?: boolean;
  /** Set to true to skip the automatic 401 → token-refresh retry. */
  _isRetry?: boolean;
};

export async function apiFetch<T>(
  path: string,
  options: RequestOptions = {}
): Promise<T> {
  const { shouldAuth = true, _isRetry = false, headers, ...rest } = options;
  const token = shouldAuth ? getAccessToken() : null;

  const response = await fetch(`${API_BASE_URL}${path}`, {
    ...rest,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
      ...(headers ?? {}),
    },
  });

  // ── Auto-refresh on 401 ───────────────────────────────────────────────────
  if (response.status === 401 && shouldAuth && !_isRetry) {
    if (!refreshPromise) {
      refreshPromise = doRefresh().finally(() => {
        refreshPromise = null;
      });
    }

    const newToken = await refreshPromise;
    if (newToken) {
      return apiFetch<T>(path, { ...options, _isRetry: true });
    }
    // No valid refresh token — caller will see the 401
  }

  if (!response.ok) {
    let message = `HTTP ${response.status}`;
    try {
      const body = await response.text();
      // Use server message only when it is a plain string (not raw HTML)
      if (body && !body.trimStart().startsWith("<")) {
        message = body;
      }
    } catch {
      // ignore
    }
    throw new ApiError(response.status, message);
  }

  if (response.status === 204) {
    return {} as T;
  }

  return (await response.json()) as T;
}
