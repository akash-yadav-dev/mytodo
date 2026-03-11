import type { FieldErrors } from "./auth";

export function validateProject(values: {
  name: string;
  description?: string;
}) {
  const errors: FieldErrors<typeof values> = {};
  if (!values.name) errors.name = "Project name is required.";
  else if (values.name.length < 2)
    errors.name = "Name must be at least 2 characters.";
  return errors;
}
