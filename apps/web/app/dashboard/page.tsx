"use client";

import { AuthGuard } from "@/components/auth/auth-guard";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { useAuthContext } from "@/components/auth/auth-provider";
import { useMyOrganizations } from "@/hooks/useOrganizations";
import { ROUTES } from "@/lib/constants/routes";
import Link from "next/link";

export default function DashboardPage() {
  return (
    <AuthGuard>
      <DashboardContent />
    </AuthGuard>
  );
}

const statCards = [
  { label: "Open issues", value: "—", trend: null, icon: "���" },
  { label: "Active projects", value: "—", trend: null, icon: "���" },
  { label: "Completed this week", value: "—", trend: null, icon: "✅" },
  { label: "Team members", value: "—", trend: null, icon: "���" },
];

const quickLinks = [
  { label: "Browse issues", href: ROUTES.issues, icon: "���" },
  { label: "View projects", href: ROUTES.projects, icon: "���" },
  { label: "Open boards", href: ROUTES.boards, icon: "���" },
  { label: "Organizations", href: ROUTES.organizations, icon: "���" },
  { label: "Profile settings", href: ROUTES.settingsProfile, icon: "���" },
];

function DashboardContent() {
  const { user } = useAuthContext();
  const { data: orgData, isLoading: orgsLoading } = useMyOrganizations();
  const orgs = orgData?.data ?? [];

  const greeting = () => {
    const h = new Date().getHours();
    if (h < 12) return "Good morning";
    if (h < 17) return "Good afternoon";
    return "Good evening";
  };

  return (
    <div className="mx-auto max-w-6xl px-6 py-10">
      {/* Welcome banner */}
      <div className="mb-8">
        <h1 className="text-2xl font-semibold">
          {greeting()}, {user?.name?.split(" ")[0] ?? "there"} ���
        </h1>
        <p className="mt-1 text-sm text-fg/60">
          Here&apos;s an overview of your workspace.
        </p>
      </div>

      {/* Stats */}
      <div className="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
        {statCards.map((stat) => (
          <Card key={stat.label} className="flex items-center gap-4 !p-5">
            <span className="text-3xl">{stat.icon}</span>
            <div>
              <p className="text-2xl font-semibold">{stat.value}</p>
              <p className="text-xs text-fg/60">{stat.label}</p>
            </div>
          </Card>
        ))}
      </div>

      <div className="grid gap-6 lg:grid-cols-3">
        {/* Quick links */}
        <div className="lg:col-span-1">
          <h2 className="mb-4 text-sm font-semibold uppercase tracking-wider text-fg/50">
            Quick links
          </h2>
          <div className="space-y-2">
            {quickLinks.map((item) => (
              <Link
                key={item.href}
                href={item.href}
                className="flex items-center gap-3 rounded-2xl border border-border bg-card px-4 py-3 text-sm transition hover:border-accent/40 hover:bg-accentMuted"
              >
                <span className="text-lg">{item.icon}</span>
                {item.label}
              </Link>
            ))}
          </div>
        </div>

        {/* Organizations */}
        <div className="lg:col-span-2">
          <div className="mb-4 flex items-center justify-between">
            <h2 className="text-sm font-semibold uppercase tracking-wider text-fg/50">
              Your organizations
            </h2>
            <Button asChild href={ROUTES.organizations} variant="ghost">
              View all
            </Button>
          </div>

          {orgsLoading && (
            <div className="space-y-3">
              {[1, 2].map((i) => (
                <div
                  key={i}
                  className="h-20 animate-pulse rounded-3xl border border-border bg-card"
                />
              ))}
            </div>
          )}

          {!orgsLoading && orgs.length === 0 && (
            <Card className="flex flex-col items-center py-10 text-center">
              <span className="mb-2 text-4xl">���</span>
              <p className="mb-3 text-sm text-fg/60">
                You haven&apos;t created any organizations yet.
              </p>
              <Button asChild href={ROUTES.organizations}>
                Create organization
              </Button>
            </Card>
          )}

          {orgs.length > 0 && (
            <div className="space-y-3">
              {orgs.slice(0, 4).map((org) => (
                <Link
                  key={org.id}
                  href={ROUTES.organizationDetail(org.id)}
                  className="flex items-center justify-between rounded-3xl border border-border bg-card px-5 py-4 shadow-soft transition hover:border-accent/40 hover:shadow-lg"
                >
                  <div>
                    <p className="font-medium">{org.name}</p>
                    <p className="text-xs text-fg/50">/{org.slug}</p>
                  </div>
                  <Badge variant={org.is_active ? "success" : "default"}>
                    {org.is_active ? "Active" : "Inactive"}
                  </Badge>
                </Link>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
