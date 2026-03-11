export type FieldErrors<T> = Partial<Record<keyof T, string>>;

export function validateLogin(values: { email: string; password: string }) {
  const errors: FieldErrors<typeof values> = {};
  if (!values.email) errors.email = "Email is required.";
  else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(values.email))
    errors.email = "Please enter a valid email address.";
  if (!values.password) errors.password = "Password is required.";
  else if (values.password.length < 6)
    errors.password = "Password must be at least 6 characters.";
  return errors;
}

export function validateRegister(values: {
  name: string;
  email: string;
  password: string;
}) {
  const errors: FieldErrors<typeof values> = {};
  if (!values.name) errors.name = "Name is required.";
  else if (values.name.length < 2)
    errors.name = "Name must be at least 2 characters.";
  const loginErrors = validateLogin(values);
  if (loginErrors.email) errors.email = loginErrors.email;
  if (loginErrors.password) errors.password = loginErrors.password;
  return errors;
}
