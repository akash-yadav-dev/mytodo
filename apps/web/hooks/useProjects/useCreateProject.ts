import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createProject } from "@/services/api/projects";
import { projectKeys } from "./useProjects";

export function useCreateProject() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: createProject,
    onSuccess: () => qc.invalidateQueries({ queryKey: projectKeys.lists() }),
  });
}
