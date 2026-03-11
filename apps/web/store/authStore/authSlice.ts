// Auth store slice — thin wrapper since auth state lives in AuthContext (react context).
// This file is a placeholder for if Zustand or Redux is added in the future.

export type AuthSliceState = {
  /** Mirrors the AuthContext user for components that prefer store access. */
  userId: string | null;
};

export const initialAuthState: AuthSliceState = {
  userId: null,
};
