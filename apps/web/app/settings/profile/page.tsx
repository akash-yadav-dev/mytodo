"use client";

import { useState } from "react";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Card } from "@/components/ui/card";
import { Avatar } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useAuthContext } from "@/components/auth/auth-provider";

export default function ProfileSettingsPage() {
  return (
    <AuthGuard>
      <ProfileContent />
    </AuthGuard>
  );
}

function ProfileContent() {
  const { user } = useAuthContext();
  const [form, setForm] = useState({
    name: user?.name ?? "",
    email: user?.email ?? "",
  });
  const [saved, setSaved] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: wire to PATCH /api/v1/users/me when backend endpoint is available
    setSaved(true);
    setTimeout(() => setSaved(false), 2000);
  };

  return (
    <div className="mx-auto max-w-2xl px-6 py-10">
      <div className="mb-8">
        <h1 className="text-2xl font-semibold">Profile settings</h1>
        <p className="mt-1 text-sm text-fg/60">
          Update your personal information.
        </p>
      </div>

      <Card className="space-y-6">
        {/* Avatar */}
        <div className="flex items-center gap-4">
          <Avatar name={user?.name ?? "?"} size="lg" />
          <div>
            <p className="font-medium">{user?.name}</p>
            <p className="text-sm text-fg/60">{user?.email}</p>
          </div>
        </div>

        <hr className="border-border" />

        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">Full name</label>
            <Input
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              placeholder="Your name"
            />
          </div>
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">Email address</label>
            <Input
              type="email"
              value={form.email}
              onChange={(e) => setForm({ ...form, email: e.target.value })}
              placeholder="you@example.com"
            />
          </div>
          <Button type="submit" className="w-full">
            {saved ? "Saved ✓" : "Save changes"}
          </Button>
        </form>
      </Card>
    </div>
  );
}
