import type { IssueSliceState } from "./issueSlice";

export const selectSelectedIssueId = (state: IssueSliceState) =>
  state.selectedIssueId;
export const selectFilterStatus = (state: IssueSliceState) =>
  state.filterStatus;
