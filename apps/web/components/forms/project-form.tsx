"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import type { Project } from "@/services/api/projects";

type CreateProjectPayload = {
  name: string;
  description?: string;
  key: string;
};

type Props = {
  initial?: Partial<Project>;
  onSubmit: (payload: CreateProjectPayload) => Promise<void>;
  onCancel?: () => void;
  loading?: boolean;
  error?: string | null;
  mode?: "create" | "edit";
};

export function ProjectForm({ initial, onSubmit, onCancel, loading, error, mode = "create" }: Props) {
  const [form, setForm] = useState<CreateProjectPayload>({
    name: initial?.name ?? "",
    description: initial?.description ?? "",
    key: initial?.key ?? "",
  });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    await onSubmit(form);
  };

  // Auto-generate key from name
  const handleNameChange = (name: string) => {
    const key = mode === "create"
      ? name.toUpperCase().replace(/[^A-Z0-9]/g, "").slice(0, 6)
      : form.key;
    setForm({ ...form, name, key });
  };

  return (
    <form className="space-y-4" onSubmit={handleSubmit}>
      <div className="space-y-1">
        <label className="text-sm font-medium">Project name *</label>
        <Input
          placeholder="e.g. Backend API, Mobile App"
          value={form.name}
          onChange={(e) => handleNameChange(e.target.value)}
          required
        />
      </div>

      <div className="space-y-1">
        <label className="text-sm font-medium">Project key *</label>
        <Input
          placeholder="e.g. API, MOB"
          value={form.key}
          onChange={(e) => setForm({ ...form, key: e.target.value.toUpperCase() })}
          maxLength={6}
          required
        />
        <p className="text-xs text-fg/50">Short prefix for issue IDs (auto-generated from name).</p>
      </div>

      <div className="space-y-1">
        <label className="text-sm font-medium">Description</label>
        <textarea
          className="w-full rounded-xl border border-border bg-card px-3 py-2 text-sm outline-none placeholder:text-fg/40 focus:border-accent/60 focus:ring-2 focus:ring-accent/20 min-h-[80px] resize-y"
          placeholder="What is this project about?"
          value={form.description}
          onChange={(e) => setForm({ ...form, description: e.target.value })}
        />
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
          {loading ? (mode === "create" ? "Creating…" : "Saving…") : (mode === "create" ? "Create project" : "Save changes")}
        </Button>
      </div>
    </form>
  );
}
