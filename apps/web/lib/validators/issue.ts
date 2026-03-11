import type { FieldErrors } from "./auth";

export function validateIssue(values: {
  title: string;
  description?: string;
  priority?: string;
}) {
  const errors: FieldErrors<typeof values> = {};
  if (!values.title) errors.title = "Title is required.";
  else if (values.title.length < 3)
    errors.title = "Title must be at least 3 characters.";
  return errors;
}
