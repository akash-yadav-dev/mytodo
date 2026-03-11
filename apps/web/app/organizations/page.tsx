"use client";

import { useState } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Avatar } from "@/components/ui/avatar";
import { Modal } from "@/components/ui/modal";
import { Input } from "@/components/ui/input";
import { useMyOrganizations, useCreateOrganization } from "@/hooks/useOrganizations";
import { ROUTES } from "@/lib/constants/routes";

export default function OrganizationsPage() {
  return (
    <AuthGuard>
      <OrganizationsContent />
    </AuthGuard>
  );
}

function OrganizationsContent() {
  const { data, isLoading, error } = useMyOrganizations();
  const { mutate: create, isPending } = useCreateOrganization();
  const [showCreate, setShowCreate] = useState(false);
  const [form, setForm] = useState({ name: "", slug: "", description: "" });

  const orgs = data?.data ?? [];

  const handleCreate = (e: React.FormEvent) => {
    e.preventDefault();
    create(
      { name: form.name, slug: form.slug || undefined, description: form.description || undefined },
      {
        onSuccess: () => {
          setShowCreate(false);
          setForm({ name: "", slug: "", description: "" });
        },
      }
    );
  };

  return (
    <div className="mx-auto max-w-6xl px-6 py-10">
      <div className="mb-8 flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Organizations</h1>
          <p className="mt-1 text-sm text-fg/60">
            Manage your teams and workspaces.
          </p>
        </div>
        <Button onClick={() => setShowCreate(true)}>New organization</Button>
      </div>

      {isLoading && (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {[1, 2, 3].map((i) => (
            <div
              key={i}
              className="h-36 animate-pulse rounded-3xl border border-border bg-card"
            />
          ))}
        </div>
      )}

      {error && (
        <div className="rounded-2xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-600">
          Failed to load organizations. Please try again.
        </div>
      )}

      {!isLoading && orgs.length === 0 && (
        <div className="flex flex-col items-center justify-center rounded-3xl border border-dashed border-border py-20 text-center">
          <span className="mb-3 text-5xl">🏢</span>
          <h3 className="mb-2 text-lg font-medium">No organizations yet</h3>
          <p className="mb-6 max-w-xs text-sm text-fg/60">
            Create your first organization to start collaborating with your team.
          </p>
          <Button onClick={() => setShowCreate(true)}>
            Create organization
          </Button>
        </div>
      )}

      {orgs.length > 0 && (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {orgs.map((org) => (
            <Link
              key={org.id}
              href={ROUTES.organizationDetail(org.id)}
              className="group rounded-3xl border border-border bg-card p-6 shadow-soft transition hover:-translate-y-1 hover:border-accent/40 hover:shadow-lg"
            >
              <div className="mb-4 flex items-start justify-between">
                <Avatar name={org.name} size="lg" />
                <Badge variant={org.is_active ? "success" : "default"}>
                  {org.is_active ? "Active" : "Inactive"}
                </Badge>
              </div>
              <h3 className="text-lg font-semibold">{org.name}</h3>
              <p className="mt-1 text-xs text-fg/50">/{org.slug}</p>
              {org.description && (
                <p className="mt-2 line-clamp-2 text-sm text-fg/70">
                  {org.description}
                </p>
              )}
              <p className="mt-4 text-xs text-fg/40">
                Created{" "}
                {new Date(org.created_at).toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                  year: "numeric",
                })}
              </p>
            </Link>
          ))}
        </div>
      )}

      {/* Create modal */}
      <Modal
        open={showCreate}
        onClose={() => setShowCreate(false)}
        title="Create organization"
      >
        <form onSubmit={handleCreate} className="space-y-4">
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">Name *</label>
            <Input
              placeholder="Acme Corp"
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              required
            />
          </div>
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">
              Slug (optional)
            </label>
            <Input
              placeholder="acme-corp"
              value={form.slug}
              onChange={(e) => setForm({ ...form, slug: e.target.value })}
            />
          </div>
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">
              Description (optional)
            </label>
            <Input
              placeholder="What does your org do?"
              value={form.description}
              onChange={(e) =>
                setForm({ ...form, description: e.target.value })
              }
            />
          </div>
          <div className="flex gap-2 pt-2">
            <Button type="submit" disabled={isPending} className="flex-1">
              {isPending ? "Creating..." : "Create"}
            </Button>
            <Button
              type="button"
              variant="secondary"
              onClick={() => setShowCreate(false)}
            >
              Cancel
            </Button>
          </div>
        </form>
      </Modal>
    </div>
  );
}
