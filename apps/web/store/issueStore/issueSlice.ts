// Issue store slice placeholder — actual state is managed via React Query.
export type IssueSliceState = {
  selectedIssueId: string | null;
  filterStatus: string | null;
};

export const initialIssueState: IssueSliceState = {
  selectedIssueId: null,
  filterStatus: null,
};
