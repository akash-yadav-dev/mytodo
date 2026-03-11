"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import type { CreateIssuePayload } from "@/services/api/issues";

type Props = {
  projectId: string;
  onSubmit: (payload: CreateIssuePayload) => Promise<void>;
  onCancel?: () => void;
  loading?: boolean;
  error?: string | null;
};

const PRIORITIES = ["low", "medium", "high", "urgent"] as const;
const STATUSES = ["open", "in_progress", "done", "cancelled"] as const;

export function IssueForm({ projectId, onSubmit, onCancel, loading, error }: Props) {
  const [form, setForm] = useState<CreateIssuePayload>({
    title: "",
    description: "",
    status: "open",
    priority: "medium",
    project_id: projectId,
    assignee_id: undefined,
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(form);
  };

  return (
    <form className="space-y-4" onSubmit={handleSubmit}>
      <div className="space-y-1">
        <label className="text-sm font-medium">Title *</label>
        <Input
          placeholder="Short, descriptive issue title"
          value={form.title}
          onChange={(e) => setForm({ ...form, title: e.target.value })}
          required
        />
      </div>

      <div className="space-y-1">
        <label className="text-sm font-medium">Description</label>
        <textarea
          className="w-full rounded-xl border border-border bg-card px-3 py-2 text-sm outline-none placeholder:text-fg/40 focus:border-accent/60 focus:ring-2 focus:ring-accent/20 min-h-[100px] resize-y"
          placeholder="Describe the issue in more detail…"
          value={form.description}
          onChange={(e) => setForm({ ...form, description: e.target.value })}
        />
      </div>

      <div className="grid grid-cols-2 gap-4">
        <div className="space-y-1">
          <label className="text-sm font-medium">Status</label>
          <select
            className="w-full rounded-xl border border-border bg-card px-3 py-2 text-sm outline-none focus:border-accent/60"
            value={form.status}
            onChange={(e) => setForm({ ...form, status: e.target.value as typeof form.status })}
          >
            {STATUSES.map((s) => (
              <option key={s} value={s}>{s.replace("_", " ")}</option>
            ))}
          </select>
        </div>

        <div className="space-y-1">
          <label className="text-sm font-medium">Priority</label>
          <select
            className="w-full rounded-xl border border-border bg-card px-3 py-2 text-sm outline-none focus:border-accent/60"
            value={form.priority}
            onChange={(e) => setForm({ ...form, priority: e.target.value as typeof form.priority })}
          >
            {PRIORITIES.map((p) => (
              <option key={p} value={p}>{p}</option>
            ))}
          </select>
        </div>
      </div>

      {error && (
        <p className="rounded-lg border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-600">
          {error}
        </p>
      )}

      <div className="flex justify-end gap-3 pt-2">
        {onCancel && (
          <Button type="button" variant="secondary" onClick={onCancel}>
            Cancel
          </Button>
        )}
        <Button type="submit" disabled={loading}>
          {loading ? "Creating…" : "Create issue"}
        </Button>
      </div>
    </form>
  );
}
