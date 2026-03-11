import type { AuthUser } from "@/lib/types/auth";

const SESSION_KEY = "mytodo_session";

type StoredSession = { user: AuthUser };

export function getStoredSession(): StoredSession | null {
  try {
    const raw = localStorage.getItem(SESSION_KEY);
    if (!raw) return null;
    return JSON.parse(raw) as StoredSession;
  } catch {
    return null;
  }
}

export function setStoredSession(session: StoredSession): void {
  try {
    localStorage.setItem(SESSION_KEY, JSON.stringify(session));
  } catch {}
}

export function clearStoredSession(): void {
  try {
    localStorage.removeItem(SESSION_KEY);
  } catch {}
}
