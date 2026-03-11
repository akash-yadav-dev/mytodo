"use client";

import Link from "next/link";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Button } from "@/components/ui/button";
import { Avatar } from "@/components/ui/avatar";
import { Dropdown } from "@/components/ui/dropdown";
import { APP_NAME } from "@/lib/constants/config";
import { ROUTES } from "@/lib/constants/routes";

const marketingNav = [
  { label: "Features", href: "/#features" },
  { label: "Pricing", href: "/#pricing" },
  { label: "Docs", href: "/docs" },
];

export function SiteHeader() {
  const { user, status, logout } = useAuthContext();
  const isAuthenticated = status === "authenticated";

  return (
    <header className="sticky top-0 z-30 border-b border-border bg-card/80 backdrop-blur supports-[backdrop-filter]:bg-card/60">
      <div className="mx-auto flex h-16 max-w-7xl items-center justify-between px-6">
        {/* Logo */}
        <Link
          href={isAuthenticated ? ROUTES.dashboard : ROUTES.home}
          className="flex items-center gap-2 font-semibold text-fg"
        >
          <span className="flex h-7 w-7 items-center justify-center rounded-lg bg-accent text-xs font-bold text-white">
            MT
          </span>
          <span className="hidden sm:inline">{APP_NAME}</span>
        </Link>

        {/* Marketing nav — only for guests */}
        {!isAuthenticated && (
          <nav className="hidden items-center gap-6 text-sm md:flex">
            {marketingNav.map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className="text-fg/70 transition hover:text-fg"
              >
                {item.label}
              </Link>
            ))}
          </nav>
        )}

        {/* App nav — for authenticated users */}
        {isAuthenticated && (
          <nav className="hidden items-center gap-1 text-sm md:flex">
            {[
              { label: "Dashboard", href: ROUTES.dashboard },
              { label: "Issues", href: ROUTES.issues },
              { label: "Projects", href: ROUTES.projects },
              { label: "Organizations", href: ROUTES.organizations },
            ].map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className="rounded-lg px-3 py-2 text-fg/70 transition hover:bg-accentMuted hover:text-fg"
              >
                {item.label}
              </Link>
            ))}
          </nav>
        )}

        {/* Right actions */}
        <div className="flex items-center gap-2">
          <ThemeToggle />

          {status === "guest" && (
            <>
              <Button asChild href={ROUTES.login} variant="ghost">
                Sign in
              </Button>
              <Button asChild href={ROUTES.register}>
                Get started
              </Button>
            </>
          )}

          {isAuthenticated && user && (
            <Dropdown
              trigger={
                <button className="flex items-center gap-2 rounded-full transition hover:opacity-80">
                  <Avatar name={user.name} size="sm" />
                  <span className="hidden text-sm font-medium sm:inline">
                    {user.name}
                  </span>
                </button>
              }
              items={[
                {
                  label: "Profile settings",
                  onClick: () => (window.location.href = ROUTES.settingsProfile),
                },
                {
                  label: "Organizations",
                  onClick: () => (window.location.href = ROUTES.organizations),
                },
                {
                  label: "Sign out",
                  onClick: logout,
                  destructive: true,
                },
              ]}
            />
          )}
        </div>
      </div>
    </header>
  );
}
