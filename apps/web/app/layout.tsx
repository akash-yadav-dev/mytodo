import type { Metadata } from "next";
import { Space_Grotesk } from "next/font/google";
import "../styles/globals.css";
import { AppProviders } from "@/components/providers/app-providers";

const spaceGrotesk = Space_Grotesk({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "MyTodo",
  description: "A clean, extensible workspace for tasks and projects.",
};



export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
  }) {
  const STORAGE_KEY = "mytodo-theme";
  
  // Move the script to a constant and make it more robust
  const themeScript = `
    (function () {
      try {
        var stored = localStorage.getItem(STORAGE_KEY);
        var prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        var isDark = stored === 'dark' || (!stored && prefersDark);
        if (isDark) document.documentElement.classList.add('dark');
      } catch (_) {}
      // Enable transitions a tick after first paint to prevent flash
      requestAnimationFrame(function () {
        requestAnimationFrame(function () {
          document.documentElement.classList.add('theme-ready');
        });
      });
    })();
    `.trim();

  return (
    <html lang="en" suppressHydrationWarning>
      <head>
        {/* Script runs before any content is painted */}
        <script dangerouslySetInnerHTML={{ __html: themeScript }} />
      </head>
      <body className={spaceGrotesk.className}>
        <AppProviders>{children}</AppProviders>
      </body>
    </html>
  );
}
