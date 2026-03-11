"use client";

import { useTheme } from "@/components/theme/theme-provider";

export function ThemeToggle() {
  const { theme, toggle } = useTheme();
  return (
    <button
      onClick={toggle}
      aria-label="Toggle theme"
      className="flex h-8 w-8 items-center justify-center rounded-lg border border-border bg-card text-fg/70 transition hover:bg-accentMuted hover:text-accent"
    >
      {theme === "dark" ? "☀" : "☾"}
    </button>
  );
}
