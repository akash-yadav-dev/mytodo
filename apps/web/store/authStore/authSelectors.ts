import type { AuthSliceState } from "./authSlice";

export const selectUserId = (state: AuthSliceState) => state.userId;
export const selectIsAuthenticated = (state: AuthSliceState) =>
  state.userId !== null;
