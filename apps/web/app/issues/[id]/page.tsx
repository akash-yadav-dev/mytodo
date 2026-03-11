"use client";

import { use } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";

export default function IssueDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);

  return (
    <AuthGuard>
      <div className="mx-auto max-w-4xl px-6 py-10">
        {/* Breadcrumb */}
        <nav className="mb-6 flex items-center gap-2 text-sm text-fg/50">
          <Link href={ROUTES.issues} className="transition hover:text-fg">
            Issues
          </Link>
          <span>/</span>
          <span className="font-mono text-fg">{id}</span>
        </nav>

        <div className="grid gap-6 lg:grid-cols-3">
          {/* Main content */}
          <div className="lg:col-span-2 space-y-4">
            <Card>
              <div className="mb-4 flex items-start justify-between gap-4">
                <h1 className="text-xl font-semibold">Issue #{id}</h1>
                <Badge variant="default">Open</Badge>
              </div>
              <p className="text-sm text-fg/70">
                Issue details will load from the backend once the issues API is
                fully wired in. This is a placeholder skeleton.
              </p>
            </Card>

            {/* Comments placeholder */}
            <Card>
              <h3 className="mb-3 font-medium">Comments</h3>
              <p className="text-sm text-fg/50">No comments yet.</p>
            </Card>
          </div>

          {/* Sidebar */}
          <aside className="space-y-4">
            <Card>
              <h3 className="mb-3 text-sm font-medium text-fg/60">Details</h3>
              <dl className="space-y-2 text-sm">
                <div className="flex justify-between">
                  <dt className="text-fg/50">Status</dt>
                  <dd>Open</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-fg/50">Priority</dt>
                  <dd>Medium</dd>
                </div>
                <div className="flex justify-between">
                  <dt className="text-fg/50">Assignee</dt>
                  <dd className="text-fg/50">Unassigned</dd>
                </div>
              </dl>
            </Card>
            <Button variant="secondary" asChild href={ROUTES.issues} className="w-full">
              Back to issues
            </Button>
          </aside>
        </div>
      </div>
    </AuthGuard>
  );
}
