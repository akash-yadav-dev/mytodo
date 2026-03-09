import { useState } from "react";
import { useAuth } from "@/hooks/useAuth/useAuth";
import type { LoginPayload } from "@/lib/types/auth";

export function useLogin() {
  const { login } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const submit = async (payload: LoginPayload) => {
    setLoading(true);
    setError(null);
    try {
      await login(payload);
      return true;
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed");
      return false;
    } finally {
      setLoading(false);
    }
  };

  return { submit, loading, error };
}
