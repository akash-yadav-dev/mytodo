import { useQuery } from "@tanstack/react-query";
import { getCurrentUser } from "@/services/api/users";

export const userKeys = {
  me: ["users", "me"] as const,
};

export function useCurrentUser() {
  return useQuery({
    queryKey: userKeys.me,
    queryFn: getCurrentUser,
  });
}
