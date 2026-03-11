import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createIssue } from "@/services/api/issues";
import { issueKeys } from "./useIssues";

export function useCreateIssue() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: createIssue,
    onSuccess: () => qc.invalidateQueries({ queryKey: issueKeys.lists() }),
  });
}
