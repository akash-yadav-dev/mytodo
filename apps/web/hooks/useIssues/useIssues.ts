import { useQuery } from "@tanstack/react-query";
import { listIssues } from "@/services/api/issues";

export const issueKeys = {
  all: ["issues"] as const,
  lists: () => [...issueKeys.all, "list"] as const,
  detail: (id: string) => [...issueKeys.all, "detail", id] as const,
};

export function useIssues(projectId?: string) {
  return useQuery({
    queryKey: [...issueKeys.lists(), { projectId }],
    queryFn: () => listIssues(projectId),
  });
}
