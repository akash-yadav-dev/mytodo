import type { Metadata } from "next";
import { Space_Grotesk } from "next/font/google";
import "../styles/globals.css";
import { AppProviders } from "@/components/providers/app-providers";
import { LayoutShell } from "@/components/layouts/layout-shell";

const spaceGrotesk = Space_Grotesk({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: {
    template: "%s | MyTodo",
    default: "MyTodo – Ticket management for productive teams",
  },
  description:
    "A clean, extensible workspace for tasks, tickets, and projects. Built for teams that move fast.",
};

const themeScript = [
  "(function () {",
  "  try {",
  "    var stored = localStorage.getItem('mytodo-theme');",
  "    var prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;",
  "    var isDark = stored === 'dark' || (!stored && prefersDark);",
  "    if (isDark) document.documentElement.setAttribute('data-theme', 'dark');",
  "    else document.documentElement.setAttribute('data-theme', 'light');",
  "  } catch (_) {}",
  "  requestAnimationFrame(function () {",
  "    requestAnimationFrame(function () {",
  "      document.documentElement.classList.add('theme-ready');",
  "    });",
  "  });",
  "})();",
].join("\n");

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        <script dangerouslySetInnerHTML={{ __html: themeScript }} />
      </head>
      <body className={spaceGrotesk.className}>
        <AppProviders>
          <LayoutShell>{children}</LayoutShell>
        </AppProviders>
      </body>
    </html>
  );
}
