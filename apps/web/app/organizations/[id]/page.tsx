"use client";

import { use, useState } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Avatar } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Modal } from "@/components/ui/modal";
import { Input } from "@/components/ui/input";
import {
  useOrganization,
  useOrganizationMembers,
  useAddMember,
  useRemoveMember,
} from "@/hooks/useOrganizations";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ROUTES } from "@/lib/constants/routes";

export default function OrganizationDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  return (
    <AuthGuard>
      <OrgContent id={id} />
    </AuthGuard>
  );
}

function OrgContent({ id }: { id: string }) {
  const { user } = useAuthContext();
  const { data: org, isLoading, error } = useOrganization(id);
  const { data: members = [], isLoading: membersLoading } = useOrganizationMembers(id);
  const addMember = useAddMember(id);
  const removeMember = useRemoveMember(id);

  const [addModalOpen, setAddModalOpen] = useState(false);
  const [inviteEmail, setInviteEmail] = useState("");
  const [inviteRole, setInviteRole] = useState("member");
  const [addError, setAddError] = useState<string | null>(null);

  const isOwner = org?.owner_id === user?.id;

  if (isLoading) {
    return (
      <div className="mx-auto max-w-4xl px-6 py-10 space-y-4">
        <div className="h-8 w-48 animate-pulse rounded-xl bg-card" />
        <div className="h-4 w-72 animate-pulse rounded-xl bg-card" />
        <div className="mt-8 grid gap-4 sm:grid-cols-3">
          {[1, 2, 3].map((n) => <div key={n} className="h-32 animate-pulse rounded-2xl bg-card" />)}
        </div>
      </div>
    );
  }

  if (error || !org) {
    return (
      <div className="mx-auto max-w-4xl px-6 py-10">
        <div className="rounded-2xl border border-red-200 bg-red-50 px-4 py-3 text-sm text-red-600">
          Organization not found or you don&apos;t have access.
        </div>
        <Button asChild href={ROUTES.organizations} variant="secondary" className="mt-4">
          Back
        </Button>
      </div>
    );
  }

  const handleAddMember = async (e: React.FormEvent) => {
    e.preventDefault();
    setAddError(null);
    try {
      await addMember.mutateAsync({ email: inviteEmail, role: inviteRole });
      setInviteEmail("");
      setAddModalOpen(false);
    } catch (err) {
      setAddError(err instanceof Error ? err.message : "Failed to add member");
    }
  };

  return (
    <div className="mx-auto max-w-4xl px-6 py-10">
      {/* Breadcrumb */}
      <nav className="mb-6 flex items-center gap-2 text-sm text-fg/50">
        <Link href={ROUTES.organizations} className="hover:text-fg">Organizations</Link>
        <span>/</span>
        <span className="text-fg">{org.name}</span>
      </nav>

      {/* Org header */}
      <div className="mb-8 flex items-start justify-between">
        <div className="flex items-center gap-4">
          <div className="flex h-14 w-14 items-center justify-center rounded-2xl bg-accent text-xl font-bold text-white">
            {org.name[0]?.toUpperCase()}
          </div>
          <div>
            <h1 className="text-2xl font-semibold">{org.name}</h1>
            {org.description && (
              <p className="mt-1 text-sm text-fg/60">{org.description}</p>
            )}
            <div className="mt-1 flex items-center gap-2">
              {org.slug && (
                <span className="font-mono text-xs text-fg/40">/{org.slug}</span>
              )}
              {isOwner && <Badge>Owner</Badge>}
            </div>
          </div>
        </div>
        {isOwner && (
          <Button variant="secondary" onClick={() => setAddModalOpen(true)}>
            Add member
          </Button>
        )}
      </div>

      {/* Stats */}
      <div className="mb-8 grid gap-4 sm:grid-cols-3">
        {[
          { label: "Members", value: members.length },
          { label: "Plan", value: org.plan ?? "Free" },
          { label: "Created", value: new Date(org.created_at).toLocaleDateString() },
        ].map((stat) => (
          <Card key={stat.label} className="text-center">
            <p className="text-2xl font-bold">{stat.value}</p>
            <p className="mt-1 text-sm text-fg/60">{stat.label}</p>
          </Card>
        ))}
      </div>

      {/* Members list */}
      <div>
        <h2 className="mb-4 text-lg font-semibold">Members</h2>
        {membersLoading ? (
          <div className="space-y-3">
            {[1, 2, 3].map((n) => (
              <div key={n} className="h-16 animate-pulse rounded-2xl bg-card" />
            ))}
          </div>
        ) : members.length === 0 ? (
          <div className="rounded-2xl border border-dashed border-border py-12 text-center">
            <p className="text-sm text-fg/50">No members yet. Add someone to get started.</p>
          </div>
        ) : (
          <div className="space-y-2">
            {members.map((member) => (
              <div
                key={member.id}
                className="flex items-center justify-between rounded-2xl border border-border bg-card px-5 py-4"
              >
                <div className="flex items-center gap-3">
                  <Avatar name={member.name ?? member.email} size="sm" />
                  <div>
                    <p className="font-medium">{member.name ?? member.email}</p>
                    <p className="text-xs text-fg/50">{member.email}</p>
                  </div>
                </div>
                <div className="flex items-center gap-3">
                  <Badge>{member.role}</Badge>
                  {isOwner && member.id !== user?.id && (
                    <button
                      className="text-xs text-red-500 hover:text-red-700"
                      onClick={() => removeMember.mutate(member.id)}
                    >
                      Remove
                    </button>
                  )}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Add member modal */}
      <Modal
        open={addModalOpen}
        onClose={() => setAddModalOpen(false)}
        title="Add member"
      >
        <form className="space-y-4" onSubmit={handleAddMember}>
          <div className="space-y-1">
            <label className="text-sm font-medium">Email address</label>
            <Input
              type="email"
              placeholder="colleague@company.com"
              value={inviteEmail}
              onChange={(e) => setInviteEmail(e.target.value)}
              required
            />
          </div>
          <div className="space-y-1">
            <label className="text-sm font-medium">Role</label>
            <select
              className="w-full rounded-xl border border-border bg-card px-3 py-2 text-sm outline-none focus:border-accent/60"
              value={inviteRole}
              onChange={(e) => setInviteRole(e.target.value)}
            >
              <option value="member">Member</option>
              <option value="admin">Admin</option>
            </select>
          </div>
          {addError && (
            <p className="rounded-lg border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-600">
              {addError}
            </p>
          )}
          <div className="flex justify-end gap-3">
            <Button type="button" variant="secondary" onClick={() => setAddModalOpen(false)}>
              Cancel
            </Button>
            <Button type="submit" disabled={addMember.isPending}>
              {addMember.isPending ? "Adding…" : "Add member"}
            </Button>
          </div>
        </form>
      </Modal>
    </div>
  );
}
