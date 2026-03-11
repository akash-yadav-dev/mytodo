/** Format a date string to a short human-readable string, e.g. "Mar 10, 2026". */
export function formatDate(iso: string): string {
  return new Date(iso).toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  });
}

/** Returns a relative time string such as "3 days ago". */
export function timeAgo(iso: string): string {
  const seconds = Math.floor((Date.now() - new Date(iso).getTime()) / 1000);
  const intervals: [number, string][] = [
    [31536000, "year"],
    [2592000, "month"],
    [86400, "day"],
    [3600, "hour"],
    [60, "minute"],
    [1, "second"],
  ];

  for (const [sec, label] of intervals) {
    const count = Math.floor(seconds / sec);
    if (count >= 1) {
      return `${count} ${label}${count !== 1 ? "s" : ""} ago`;
    }
  }

  return "just now";
}
