import { useAuth } from "@/hooks/useAuth/useAuth";

export function useLogout() {
  const { logout } = useAuth();
  return { logout };
}
