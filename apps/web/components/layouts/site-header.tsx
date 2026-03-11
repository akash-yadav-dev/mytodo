"use client";

import { useState, useEffect } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Menu, X } from "lucide-react";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Button } from "@/components/ui/button";
import { Avatar } from "@/components/ui/avatar";
import { Dropdown } from "@/components/ui/dropdown";
import { APP_NAME } from "@/lib/constants/config";
import { ROUTES } from "@/lib/constants/routes";
import { cn } from "@/lib/utils/cn";

const marketingNav = [
  { label: "Features", href: "/#features" },
  { label: "Pricing", href: "/#pricing" },
  { label: "Docs", href: ROUTES.docs },
];

const appNav = [
  { label: "Dashboard", href: ROUTES.dashboard },
  { label: "Issues", href: ROUTES.issues },
  { label: "Projects", href: ROUTES.projects },
  { label: "Organizations", href: ROUTES.organizations },
];

export function SiteHeader() {
  const { user, status, logout } = useAuthContext();
  const isAuthenticated = status === "authenticated";
  const pathname = usePathname();
  const [mobileOpen, setMobileOpen] = useState(false);

  // Close mobile nav on route change
  useEffect(() => {
    setMobileOpen(false);
  }, [pathname]);

  // Prevent body scroll when mobile menu is open
  useEffect(() => {
    if (mobileOpen) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "";
    }
    return () => {
      document.body.style.overflow = "";
    };
  }, [mobileOpen]);

  const navItems = isAuthenticated ? appNav : marketingNav;

  return (
    <>
      <header className="sticky top-0 z-40 border-b border-border bg-card/80 backdrop-blur supports-[backdrop-filter]:bg-card/60">
        <div className="relative mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6">
          {/* Logo */}
          <Link
            href={isAuthenticated ? ROUTES.dashboard : ROUTES.home}
            className="flex shrink-0 items-center gap-2 font-semibold text-fg"
            aria-label={`${APP_NAME} home`}
          >
            <span className="flex h-7 w-7 shrink-0 items-center justify-center rounded-lg bg-accent text-xs font-bold text-white">
              MT
            </span>
            <span className="hidden sm:inline">{APP_NAME}</span>
          </Link>

          {/* Desktop nav — absolutely centred so it never squeezes logo or right-actions */}
          <nav
            className="absolute left-1/2 hidden -translate-x-1/2 items-center gap-1 text-sm md:flex"
            aria-label="Primary navigation"
          >
            {navItems.map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className={cn(
                  "rounded-lg px-3 py-2 transition",
                  isAuthenticated
                    ? "text-fg/70 hover:bg-accentMuted hover:text-fg"
                    : "text-fg/70 hover:text-fg",
                  pathname === item.href && "text-fg font-medium"
                )}
              >
                {item.label}
              </Link>
            ))}
          </nav>

          {/* Right actions */}
          <div className="flex shrink-0 items-center gap-2">
            <ThemeToggle />

            {status === "guest" && (
              <>
                <Button asChild href={ROUTES.login} variant="ghost" className="hidden sm:inline-flex">
                  Sign in
                </Button>
                <Button asChild href={ROUTES.register} className="hidden sm:inline-flex">
                  Get started
                </Button>
              </>
            )}

            {/* Authenticated — loading skeleton while user object is still fetched */}
            {isAuthenticated && !user && (
              <div className="h-7 w-7 shrink-0 animate-pulse rounded-full bg-accentMuted" aria-label="Loading account" />
            )}

            {isAuthenticated && user && (
              <Dropdown
                trigger={
                  <button
                    className="flex items-center gap-2 rounded-full transition hover:opacity-80"
                    aria-label="Account menu"
                  >
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

            {/* Hamburger — mobile only */}
            <button
              className="flex h-9 w-9 items-center justify-center rounded-lg text-fg/70 transition hover:bg-accentMuted hover:text-fg md:hidden"
              aria-label={mobileOpen ? "Close menu" : "Open menu"}
              aria-expanded={mobileOpen}
              aria-controls="mobile-nav"
              onClick={() => setMobileOpen((v) => !v)}
            >
              {mobileOpen ? <X size={20} /> : <Menu size={20} />}
            </button>
          </div>
        </div>
      </header>

      {/* Mobile nav overlay */}
      {mobileOpen && (
        <div
          className="fixed inset-0 z-30 bg-black/40 md:hidden"
          aria-hidden="true"
          onClick={() => setMobileOpen(false)}
        />
      )}

      {/* Mobile nav drawer */}
      <div
        id="mobile-nav"
        className={cn(
          "fixed inset-x-0 top-16 z-30 border-b border-border bg-card shadow-lg transition-all duration-200 ease-in-out md:hidden",
          mobileOpen ? "translate-y-0 opacity-100" : "-translate-y-4 pointer-events-none opacity-0"
        )}
        aria-hidden={!mobileOpen}
      >
        <nav className="mx-auto max-w-7xl px-4 py-4" aria-label="Mobile navigation">
          <ul className="flex flex-col gap-1">
            {navItems.map((item) => (
              <li key={item.href}>
                <Link
                  href={item.href}
                  className={cn(
                    "block rounded-xl px-4 py-3 text-sm font-medium transition hover:bg-accentMuted hover:text-fg",
                    pathname === item.href
                      ? "bg-accentMuted text-fg"
                      : "text-fg/70"
                  )}
                >
                  {item.label}
                </Link>
              </li>
            ))}
          </ul>

          {/* Auth actions in mobile menu for guests */}
          {status === "guest" && (
            <div className="mt-4 flex flex-col gap-2 border-t border-border pt-4">
              <Button asChild href={ROUTES.login} variant="secondary" className="w-full">
                Sign in
              </Button>
              <Button asChild href={ROUTES.register} className="w-full">
                Get started free
              </Button>
            </div>
          )}
        </nav>
      </div>
    </>
  );
}
