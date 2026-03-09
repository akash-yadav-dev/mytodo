import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        bg: "var(--color-bg)",
        fg: "var(--color-fg)",
        card: "var(--color-card)",
        border: "var(--color-border)",
        accent: "var(--color-accent)",
        accentMuted: "var(--color-accent-muted)",
      },
      boxShadow: {
        soft: "0 10px 30px -12px rgba(0,0,0,0.25)",
      },
    },
  },
  plugins: [],
};

export default config;
