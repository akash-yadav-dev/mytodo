"use client";

import { usePathname } from "next/navigation";
import { SiteHeader } from "@/components/layouts/site-header";
import { SiteFooter } from "@/components/layouts/site-footer";

/**
 * Wraps every page with the correct chrome:
 *  - /auth/*        → bare children (AuthLayout handles its own UI)
 *  - /dashboard/*, /issues/*, /projects/*, /boards/*, /settings/*, /organizations/*
 *                   → SiteHeader only (MainLayout adds sidebar/content wrapper)
 *  - /              → SiteHeader + children + SiteFooter (marketing landing)
 */
export function LayoutShell({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();

  const isAuth = pathname.startsWith("/auth");
  const isApp =
    pathname.startsWith("/dashboard") ||
    pathname.startsWith("/issues") ||
    pathname.startsWith("/projects") ||
    pathname.startsWith("/boards") ||
    pathname.startsWith("/settings") ||
    pathname.startsWith("/organizations");

  if (isAuth) {
    return <>{children}</>;
  }

  if (isApp) {
    // App pages use their own MainLayout, but we still provide the
    // consistent sticky SiteHeader at the very top.
    return (
      <>
        <SiteHeader />
        <div className="min-h-[calc(100vh-4rem)]">{children}</div>
      </>
    );
  }

  // Marketing / public pages
  return (
    <div className="flex min-h-screen flex-col">
      <SiteHeader />
      <main className="flex-1">{children}</main>
      <SiteFooter />
    </div>
  );
}
