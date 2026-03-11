"use client";

type Filters = {
  status: string;
  priority: string;
};

type Props = {
  filters: Filters;
  onChange: (filters: Filters) => void;
};

const STATUSES = ["all", "open", "in_progress", "done", "cancelled"];
const PRIORITIES = ["all", "urgent", "high", "medium", "low"];

export function IssueFilters({ filters, onChange }: Props) {
  return (
    <div className="flex flex-wrap items-center gap-3">
      <div className="flex items-center gap-1.5">
        <label className="text-xs font-medium text-fg/50">Status</label>
        <select
          className="rounded-lg border border-border bg-card px-2 py-1 text-sm outline-none focus:border-accent/60"
          value={filters.status}
          onChange={(e) => onChange({ ...filters, status: e.target.value })}
        >
          {STATUSES.map((s) => (
            <option key={s} value={s}>{s === "all" ? "All statuses" : s.replace("_", " ")}</option>
          ))}
        </select>
      </div>

      <div className="flex items-center gap-1.5">
        <label className="text-xs font-medium text-fg/50">Priority</label>
        <select
          className="rounded-lg border border-border bg-card px-2 py-1 text-sm outline-none focus:border-accent/60"
          value={filters.priority}
          onChange={(e) => onChange({ ...filters, priority: e.target.value })}
        >
          {PRIORITIES.map((p) => (
            <option key={p} value={p}>{p === "all" ? "All priorities" : p}</option>
          ))}
        </select>
      </div>

      {(filters.status !== "all" || filters.priority !== "all") && (
        <button
          className="text-xs text-fg/50 underline hover:text-fg"
          onClick={() => onChange({ status: "all", priority: "all" })}
        >
          Clear filters
        </button>
      )}
    </div>
  );
}
