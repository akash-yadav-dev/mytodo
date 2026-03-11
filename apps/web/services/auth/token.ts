const ACCESS_KEY = "mytodo_access_token";
const REFRESH_KEY = "mytodo_refresh_token";

export function getAccessToken(): string | null {
  try {
    return localStorage.getItem(ACCESS_KEY);
  } catch {
    return null;
  }
}

export function setAccessToken(token: string): void {
  try {
    localStorage.setItem(ACCESS_KEY, token);
  } catch {}
}

export function clearAccessToken(): void {
  try {
    localStorage.removeItem(ACCESS_KEY);
  } catch {}
}

export function getRefreshToken(): string | null {
  try {
    return localStorage.getItem(REFRESH_KEY);
  } catch {
    return null;
  }
}

export function setRefreshToken(token: string): void {
  try {
    localStorage.setItem(REFRESH_KEY, token);
  } catch {}
}

export function clearRefreshToken(): void {
  try {
    localStorage.removeItem(REFRESH_KEY);
  } catch {}
}
