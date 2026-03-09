import { useState } from "react";
import { useAuth } from "@/hooks/useAuth/useAuth";
import type { RegisterPayload } from "@/lib/types/auth";

export function useRegister() {
  const { register } = useAuth();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const submit = async (payload: RegisterPayload) => {
    setLoading(true);
    setError(null);
    try {
      await register(payload);
      return true;
    } catch (err) {
      setError(err instanceof Error ? err.message : "Registration failed");
      return false;
    } finally {
      setLoading(false);
    }
  };

  return { submit, loading, error };
}
