"use client";

import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Card } from "@/components/ui/card";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ROUTES } from "@/lib/constants/routes";

const settingsSections = [
  {
    title: "Profile",
    description: "Edit your name, avatar, and public info.",
    href: ROUTES.settingsProfile,
  },
  {
    title: "Teams & Organizations",
    description: "Manage the organizations you belong to.",
    href: ROUTES.settingsTeam,
  },
];

export default function SettingsPage() {
  const { user } = useAuthContext();

  return (
    <AuthGuard>
      <div className="mx-auto max-w-3xl px-6 py-10">
        <div className="mb-8">
          <h1 className="text-2xl font-semibold">Settings</h1>
          <p className="mt-1 text-sm text-fg/60">
            Manage your account preferences and workspace.
          </p>
        </div>

        {/* Profile summary */}
        {user && (
          <Card className="mb-8 flex items-center gap-4 p-5">
            <div className="flex h-14 w-14 items-center justify-center rounded-full bg-accent text-lg font-bold text-white">
              {user.name?.[0]?.toUpperCase() ?? "U"}
            </div>
            <div>
              <p className="font-semibold">{user.name}</p>
              <p className="text-sm text-fg/60">{user.email}</p>
            </div>
          </Card>
        )}

        {/* Settings sections */}
        <div className="grid gap-4 sm:grid-cols-2">
          {settingsSections.map((s) => (
            <Link
              key={s.href}
              href={s.href}
              className="group rounded-2xl border border-border bg-card p-5 transition hover:border-accent/40 hover:shadow-lg"
            >
              <h2 className="font-semibold group-hover:text-accent">{s.title}</h2>
              <p className="mt-1 text-sm text-fg/60">{s.description}</p>
            </Link>
          ))}

          <div className="rounded-2xl border border-border bg-card p-5">
            <h2 className="font-semibold">Theme</h2>
            <p className="mt-1 text-sm text-fg/60">
              Choose light or dark workspace.
            </p>
            <div className="mt-4">
              <ThemeToggle />
            </div>
          </div>
        </div>
      </div>
    </AuthGuard>
  );
}
