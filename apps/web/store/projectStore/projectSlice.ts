// Project store slice placeholder — actual state is managed via React Query.
export type ProjectSliceState = {
  selectedProjectId: string | null;
};

export const initialProjectState: ProjectSliceState = {
  selectedProjectId: null,
};
