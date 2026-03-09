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
  return (
    <html lang="en" data-theme="light">
      <body className={spaceGrotesk.className}>
        <AppProviders>{children}</AppProviders>
      </body>
    </html>
  );
}
