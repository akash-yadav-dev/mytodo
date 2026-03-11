import Link from "next/link";
import { IssueCard } from "@/components/issues/issue-card";
import type { Issue } from "@/services/api/issues";

type Props = {
  issues: Issue[];
  loading?: boolean;
  emptyMessage?: string;
};

const SKELETON_COUNT = 6;

export function IssueList({ issues, loading, emptyMessage = "No issues yet." }: Props) {
  if (loading) {
    return (
      <div className="space-y-2">
        {Array.from({ length: SKELETON_COUNT }).map((_, i) => (
          <div key={i} className="h-16 animate-pulse rounded-2xl bg-card" />
        ))}
      </div>
    );
  }

  if (issues.length === 0) {
    return (
      <div className="rounded-2xl border border-dashed border-border px-6 py-16 text-center">
        <p className="text-sm text-fg/50">{emptyMessage}</p>
      </div>
    );
  }

  return (
    <div className="space-y-2">
      {issues.map((issue) => (
        <IssueCard key={issue.id} issue={issue} />
      ))}
    </div>
  );
}
