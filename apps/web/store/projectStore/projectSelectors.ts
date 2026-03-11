import type { ProjectSliceState } from "./projectSlice";

export const selectSelectedProjectId = (state: ProjectSliceState) =>
  state.selectedProjectId;
