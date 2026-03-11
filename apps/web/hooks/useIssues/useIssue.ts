import { useQuery } from "@tanstack/react-query";
import { getIssue } from "@/services/api/issues";
import { issueKeys } from "./useIssues";

export function useIssue(id: string) {
  return useQuery({
    queryKey: issueKeys.detail(id),
    queryFn: () => getIssue(id),
    enabled: !!id,
  });
}
