"use client";

import { AuthGuard } from "@/components/auth/auth-guard";
import { MainLayout } from "@/components/layouts/main-layout";
import { ThemeToggle } from "@/components/theme/theme-toggle";
import { Card } from "@/components/ui/card";
import { useAuth } from "@/hooks/useAuth/useAuth";

export default function SettingsPage() {
  const { user } = useAuth();

  return (
    <AuthGuard>
      <MainLayout>
        <div className="grid gap-6">
          <div>
            <h1 className="text-2xl font-semibold">Settings</h1>
            <p className="text-sm text-fg/70">Personalize how MyTodo feels.</p>
          </div>
          <div className="grid gap-4 md:grid-cols-2">
            <Card>
              <h3 className="text-lg font-medium">Theme</h3>
              <p className="mt-2 text-sm text-fg/70">
                Choose a light or dark workspace feel.
              </p>
              <div className="mt-4">
                <ThemeToggle />
              </div>
            </Card>
            <Card>
              <h3 className="text-lg font-medium">Profile</h3>
              <p className="mt-2 text-sm text-fg/70">
                {user ? `${user.name} · ${user.email}` : "No profile loaded"}
              </p>
              <p className="mt-4 text-xs text-fg/60">
                Connect this to user settings once the API is ready.
              </p>
            </Card>
          </div>
        </div>
      </MainLayout>
    </AuthGuard>
  );
}
