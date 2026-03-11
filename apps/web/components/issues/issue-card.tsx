import Link from "next/link";
import { Badge } from "@/components/ui/badge";
import { ROUTES } from "@/lib/constants/routes";
import type { Issue } from "@/services/api/issues";

const PRIORITY_COLORS: Record<string, string> = {
  urgent: "text-red-600 bg-red-50 border-red-200",
  high: "text-orange-600 bg-orange-50 border-orange-200",
  medium: "text-amber-600 bg-amber-50 border-amber-200",
  low: "text-sky-600 bg-sky-50 border-sky-200",
};

const STATUS_COLORS: Record<string, string> = {
  open: "text-green-600 bg-green-50 border-green-200",
  in_progress: "text-amber-600 bg-amber-50 border-amber-200",
  done: "text-fg/40 bg-card border-border line-through",
  cancelled: "text-fg/30 bg-card border-border line-through",
};

type Props = {
  issue: Issue;
  compact?: boolean;
};

export function IssueCard({ issue, compact = false }: Props) {
  return (
    <Link
      href={ROUTES.issueDetail(issue.id)}
      className="flex items-center justify-between rounded-2xl border border-border bg-card px-5 py-4 shadow-soft transition hover:border-accent/40 hover:shadow-lg"
    >
      <div className="flex min-w-0 items-center gap-4">
        <span className="shrink-0 font-mono text-xs text-fg/40">{issue.id.slice(0, 8)}</span>
        <span className="truncate font-medium">{issue.title}</span>
      </div>

      {!compact && (
        <div className="ml-4 flex shrink-0 items-center gap-2">
          <span
            className={`rounded-full border px-2.5 py-0.5 text-xs font-medium ${STATUS_COLORS[issue.status] ?? ""}`}
          >
            {issue.status.replace("_", " ")}
          </span>
          <span
            className={`rounded-full border px-2.5 py-0.5 text-xs font-medium ${PRIORITY_COLORS[issue.priority] ?? ""}`}
          >
            {issue.priority}
          </span>
        </div>
      )}
    </Link>
  );
}
