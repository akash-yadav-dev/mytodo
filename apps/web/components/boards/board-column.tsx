"use client";

import type { BoardColumn } from "@/services/api/boards";

type ColumnIssue = { id: string; title: string; priority?: string };

const PRIORITY_DOT: Record<string, string> = {
  urgent: "bg-red-500",
  high: "bg-orange-400",
  medium: "bg-amber-400",
  low: "bg-sky-400",
};

type Props = {
  column: BoardColumn;
  issues?: ColumnIssue[];
};

export function BoardColumnCard({ column, issues = [] }: Props) {
  return (
    <div className="flex min-w-[280px] flex-col rounded-2xl border border-border bg-card">
      {/* Column header */}
      <div className="flex items-center justify-between border-b border-border px-4 py-3">
        <h3 className="text-sm font-semibold">{column.name}</h3>
        <span className="rounded-full bg-accentMuted px-2 py-0.5 text-xs font-medium text-accent">
          {issues.length}
        </span>
      </div>

      {/* Issue cards */}
      <div className="flex flex-1 flex-col gap-2 p-3">
        {issues.map((issue) => (
          <div
            key={issue.id}
            className="flex items-start gap-2 rounded-xl border border-border bg-bg p-3 shadow-soft transition hover:border-accent/30 hover:shadow-md"
          >
            {issue.priority && (
              <span
                className={`mt-1.5 h-2 w-2 shrink-0 rounded-full ${PRIORITY_DOT[issue.priority] ?? "bg-fg/20"}`}
              />
            )}
            <p className="text-sm leading-snug">{issue.title}</p>
          </div>
        ))}

        {issues.length === 0 && (
          <div className="flex flex-1 items-center justify-center rounded-xl border border-dashed border-border py-8">
            <p className="text-xs text-fg/40">No issues</p>
          </div>
        )}
      </div>
    </div>
  );
}
