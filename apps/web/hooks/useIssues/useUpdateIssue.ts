import { useMutation, useQueryClient } from "@tanstack/react-query";
import { updateIssue } from "@/services/api/issues";
import { issueKeys } from "./useIssues";

export function useUpdateIssue(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: Parameters<typeof updateIssue>[1]) => updateIssue(id, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: issueKeys.detail(id) });
      qc.invalidateQueries({ queryKey: issueKeys.lists() });
    },
  });
}
