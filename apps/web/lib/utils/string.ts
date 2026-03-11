/** Convert a string to a URL-friendly slug, e.g. "Hello World" → "hello-world". */
export function slugify(str: string): string {
  return str
    .toLowerCase()
    .trim()
    .replace(/[^\w\s-]/g, "")
    .replace(/[\s_-]+/g, "-")
    .replace(/^-+|-+$/g, "");
}

/** Return initials for a name, e.g. "John Doe" → "JD". */
export function getInitials(name: string, maxChars = 2): string {
  return name
    .split(" ")
    .map((w) => w[0])
    .slice(0, maxChars)
    .join("")
    .toUpperCase();
}
