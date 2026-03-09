export type AuthUser = {
  id: string;
  name: string;
  email: string;
};

export type AuthSession = {
  user: AuthUser;
  accessToken: string;
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
