"use client";

import React, { createContext, useContext, useEffect, useMemo, useState } from "react";
import type { AuthSession, AuthUser, LoginPayload, RegisterPayload } from "@/lib/types/auth";
import {
  getAccessToken,
  setAccessToken,
  clearAccessToken,
  setRefreshToken,
  clearRefreshToken,
} from "@/services/auth/token";
import { getStoredSession, setStoredSession, clearStoredSession } from "@/services/auth/session";
import { fetchCurrentUser, loginUser, registerUser, logoutUser } from "@/services/api/auth";

type AuthStatus = "loading" | "authenticated" | "guest";

type AuthContextValue = {
  user: AuthUser | null;
  status: AuthStatus;
  login: (payload: LoginPayload) => Promise<void>;
  register: (payload: RegisterPayload) => Promise<void>;
  logout: () => Promise<void>;
};

const AuthContext = createContext<AuthContextValue | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [user, setUser] = useState<AuthUser | null>(null);
  const [status, setStatus] = useState<AuthStatus>("loading");

  useEffect(() => {
    const bootstrap = async () => {
      const cached = getStoredSession();
      if (cached?.user && getAccessToken()) {
        setUser(cached.user);
      }

      try {
        const token = getAccessToken();
        if (!token) {
          setStatus("guest");
          return;
        }
        const current = await fetchCurrentUser();
        if (current) {
          setUser(current);
          setStoredSession({ user: current });
          setStatus("authenticated");
          return;
        }
      } catch {
        clearStoredSession();
        clearAccessToken();
      }

      setUser(null);
      setStatus("guest");
    };

    bootstrap();
  }, []);

  const login = async (payload: LoginPayload) => {
    setStatus("loading");
    const session = await loginUser(payload);
    applySession(session);
  };

  const register = async (payload: RegisterPayload) => {
    setStatus("loading");
    const session = await registerUser(payload);
    applySession(session);
  };

  const logout = async () => {
    setStatus("loading");
    try {
        await logoutUser();
      } catch {
        // Ignore API logout errors and clear client session.
      } finally {
        clearStoredSession();
        clearAccessToken();
        clearRefreshToken();
        setUser(null);
        setStatus("guest");
      }
  };

  const applySession = (session: AuthSession) => {
    setAccessToken(session.accessToken);
    if (session.refreshToken) setRefreshToken(session.refreshToken);
    setStoredSession({ user: session.user });
    setUser(session.user);
    setStatus("authenticated");
  };

  const value = useMemo(
    () => ({ user, status, login, register, logout }),
    [user, status]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export function useAuthContext() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuthContext must be used within AuthProvider");
  }
  return context;
}
