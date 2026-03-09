"use client";

import { Button } from "@/components/ui/button";
import { useTheme } from "@/components/theme/theme-provider";

export function ThemeToggle() {
  const { theme, setTheme } = useTheme();
  const nextTheme = theme === "light" ? "dark" : "light";

  return (
    <Button variant="ghost" onClick={() => setTheme(nextTheme)}>
      {theme === "light" ? "Dark mode" : "Light mode"}
    </Button>
  );
}
