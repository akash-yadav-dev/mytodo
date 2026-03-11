"use client";

import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { Avatar } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { useMyOrganizations } from "@/hooks/useOrganizations";
import { ROUTES } from "@/lib/constants/routes";
import Link from "next/link";

export default function TeamSettingsPage() {
  return (
    <AuthGuard>
      <TeamContent />
    </AuthGuard>
  );
}

function TeamContent() {
  const { data, isLoading } = useMyOrganizations();
  const orgs = data?.data ?? [];

  return (
    <div className="mx-auto max-w-4xl px-6 py-10">
      <div className="mb-8 flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Team & organizations</h1>
          <p className="mt-1 text-sm text-fg/60">
            Manage your organizations and team members.
          </p>
        </div>
        <Button asChild href={ROUTES.organizations}>
          New organization
        </Button>
      </div>

      {isLoading && (
        <div className="space-y-4">
          {[1, 2].map((i) => (
            <div
              key={i}
              className="h-24 animate-pulse rounded-3xl border border-border bg-card"
            />
          ))}
        </div>
      )}

      {!isLoading && orgs.length === 0 && (
        <div className="rounded-3xl border border-dashed border-border py-16 text-center">
          <span className="mb-3 block text-5xl">���</span>
          <p className="mb-4 text-sm text-fg/60">
            You haven&apos;t created any organizations yet.
          </p>
          <Button asChild href={ROUTES.organizations}>
            Create organization
          </Button>
        </div>
      )}

      {orgs.length > 0 && (
        <div className="space-y-4">
          {orgs.map((org) => (
            <div
              key={org.id}
              className="flex items-center justify-between rounded-3xl border border-border bg-card px-6 py-5 shadow-soft"
            >
              <div className="flex items-center gap-4">
                <Avatar name={org.name} size="md" />
                <div>
                  <p className="font-semibold">{org.name}</p>
                  <p className="text-xs text-fg/50">/{org.slug}</p>
                </div>
              </div>
              <div className="flex items-center gap-3">
                <Badge variant={org.is_active ? "success" : "default"}>
                  {org.is_active ? "Active" : "Inactive"}
                </Badge>
                <Link
                  href={ROUTES.organizationDetail(org.id)}
                  className="text-sm text-fg/60 transition hover:text-fg"
                >
                  Manage →
                </Link>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
