import type { Issue } from "@/services/api/issues";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";

const PRIORITY_COLORS: Record<string, string> = {
  urgent: "text-red-600 bg-red-50 border-red-200",
  high: "text-orange-600 bg-orange-50 border-orange-200",
  medium: "text-amber-600 bg-amber-50 border-amber-200",
  low: "text-sky-600 bg-sky-50 border-sky-200",
};

const STATUS_COLORS: Record<string, string> = {
  open: "text-green-600 bg-green-50 border-green-200",
  in_progress: "text-amber-600 bg-amber-50 border-amber-200",
  done: "text-fg/40 bg-card border-border",
  cancelled: "text-fg/30 bg-card border-border",
};

type Props = { issue: Issue };

export function IssueDetail({ issue }: Props) {
  return (
    <div className="grid gap-8 lg:grid-cols-3">
      {/* Main content */}
      <div className="lg:col-span-2 space-y-6">
        <h1 className="text-2xl font-semibold">{issue.title}</h1>
        <Card>
          <h2 className="mb-3 text-sm font-medium text-fg/60">Description</h2>
          {issue.description ? (
            <p className="whitespace-pre-wrap text-sm leading-relaxed">{issue.description}</p>
          ) : (
            <p className="text-sm text-fg/40 italic">No description provided.</p>
          )}
        </Card>
      </div>

      {/* Sidebar */}
      <div className="space-y-4">
        <Card>
          <h2 className="mb-4 text-sm font-semibold uppercase tracking-wide text-fg/50">Details</h2>
          <dl className="space-y-3 text-sm">
            <div className="flex items-center justify-between">
              <dt className="text-fg/60">Status</dt>
              <dd>
                <span className={`rounded-full border px-2.5 py-0.5 text-xs font-medium ${STATUS_COLORS[issue.status] ?? ""}`}>
                  {issue.status.replace("_", " ")}
                </span>
              </dd>
            </div>
            <div className="flex items-center justify-between">
              <dt className="text-fg/60">Priority</dt>
              <dd>
                <span className={`rounded-full border px-2.5 py-0.5 text-xs font-medium ${PRIORITY_COLORS[issue.priority] ?? ""}`}>
                  {issue.priority}
                </span>
              </dd>
            </div>
            <div className="flex items-center justify-between">
              <dt className="text-fg/60">ID</dt>
              <dd className="font-mono text-xs text-fg/50">{issue.id}</dd>
            </div>
            <div className="flex items-center justify-between">
              <dt className="text-fg/60">Created</dt>
              <dd className="text-fg/70">{new Date(issue.created_at).toLocaleDateString()}</dd>
            </div>
            <div className="flex items-center justify-between">
              <dt className="text-fg/60">Updated</dt>
              <dd className="text-fg/70">{new Date(issue.updated_at).toLocaleDateString()}</dd>
            </div>
          </dl>
        </Card>
      </div>
    </div>
  );
}
