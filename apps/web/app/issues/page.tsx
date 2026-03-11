"use client";

import { useState } from "react";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { Modal } from "@/components/ui/modal";
import { IssueList } from "@/components/issues/issue-list";
import { IssueFilters } from "@/components/issues/issue-filters";
import { IssueForm } from "@/components/forms/issue-form";
import { useIssues } from "@/hooks/useIssues/useIssues";
import { useCreateIssue } from "@/hooks/useIssues/useCreateIssue";
import type { CreateIssuePayload } from "@/services/api/issues";

export default function IssuesPage() {
  return (
    <AuthGuard>
      <IssuesContent />
    </AuthGuard>
  );
}

function IssuesContent() {
  const [filters, setFilters] = useState({ status: "all", priority: "all" });
  const [createOpen, setCreateOpen] = useState(false);
  const [createError, setCreateError] = useState<string | null>(null);

  const { data, isLoading } = useIssues();
  const createIssue = useCreateIssue();
  const issues = data?.data ?? [];

  const filtered = issues.filter((issue) => {
    if (filters.status !== "all" && issue.status !== filters.status) return false;
    if (filters.priority !== "all" && issue.priority !== filters.priority) return false;
    return true;
  });

  const handleCreate = async (payload: CreateIssuePayload) => {
    setCreateError(null);
    try {
      await createIssue.mutateAsync(payload);
      setCreateOpen(false);
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : "Failed to create issue");
    }
  };

  return (
    <div className="mx-auto max-w-6xl px-6 py-10">
      <div className="mb-8 flex items-start justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Issues</h1>
          <p className="mt-1 text-sm text-fg/60">Track and resolve all work items.</p>
        </div>
        <Button onClick={() => setCreateOpen(true)}>+ New issue</Button>
      </div>

      <div className="mb-6">
        <IssueFilters filters={filters} onChange={setFilters} />
      </div>

      <IssueList
        issues={filtered}
        loading={isLoading}
        emptyMessage={
          filters.status !== "all" || filters.priority !== "all"
            ? "No issues match your filters."
            : "No issues yet. Create your first issue to get started."
        }
      />

      <Modal open={createOpen} onClose={() => setCreateOpen(false)} title="New issue">
        <IssueForm
          projectId=""
          onSubmit={handleCreate}
          onCancel={() => setCreateOpen(false)}
          loading={createIssue.isPending}
          error={createError}
        />
      </Modal>
    </div>
  );
}
