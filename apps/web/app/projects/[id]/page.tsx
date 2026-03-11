"use client";

import { use } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";

export default function ProjectDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);

  return (
    <AuthGuard>
      <div className="mx-auto max-w-5xl px-6 py-10">
        <nav className="mb-6 flex items-center gap-2 text-sm text-fg/50">
          <Link href={ROUTES.projects} className="transition hover:text-fg">
            Projects
          </Link>
          <span>/</span>
          <span className="text-fg">Project #{id}</span>
        </nav>

        <div className="mb-6 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-semibold">Project #{id}</h1>
            <p className="text-sm text-fg/60">Overview and issues</p>
          </div>
          <div className="flex gap-2">
            <Button
              asChild
              href={ROUTES.projectEdit(id)}
              variant="secondary"
            >
              Edit
            </Button>
            <Button>New issue</Button>
          </div>
        </div>

        <div className="grid gap-4 sm:grid-cols-3 mb-6">
          {[
            { label: "Open", value: "—" },
            { label: "In Progress", value: "—" },
            { label: "Done", value: "—" },
          ].map((stat) => (
            <Card key={stat.label} className="text-center !p-5">
              <p className="text-3xl font-semibold">{stat.value}</p>
              <p className="text-xs text-fg/60">{stat.label}</p>
            </Card>
          ))}
        </div>

        <Card>
          <h3 className="mb-4 font-medium">Issues</h3>
          <p className="text-sm text-fg/50">
            No issues yet for this project. Create one to get started.
          </p>
        </Card>
      </div>
    </AuthGuard>
  );
}
