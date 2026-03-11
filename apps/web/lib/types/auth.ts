export type AuthUser = {
  id: string;
  email: string;
  name: string;
};

/** Raw shape returned by the backend (snake_case) */
export type AuthApiResponse = {
  access_token: string;
  refresh_token?: string;
  token_type?: string;
  expires_in?: number;
  user: AuthUser;
};

/** Normalized client-side session (camelCase) */
export type AuthSession = {
  accessToken: string;
  refreshToken?: string;
  user: AuthUser;
};

export type LoginPayload = {
  email: string;
  password: string;
};

export type RegisterPayload = {
  name: string;
  email: string;
  password: string;
};
