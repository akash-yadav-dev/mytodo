import { apiFetch } from "@/services/api/client";

export type Issue = {
  id: string;
  title: string;
  description: string;
  status: "open" | "in_progress" | "done" | "cancelled";
  priority: "low" | "medium" | "high" | "urgent";
  project_id: string;
  assignee_id?: string;
  reporter_id: string;
  created_at: string;
  updated_at: string;
};

export type IssueListResponse = {
  data: Issue[];
  total: number;
  page: number;
  limit: number;
};

export type CreateIssuePayload = {
  title: string;
  description?: string;
  status?: Issue["status"];
  priority?: Issue["priority"];
  project_id: string;
  assignee_id?: string;
};

export type UpdateIssuePayload = Partial<
  Omit<CreateIssuePayload, "project_id">
>;

export async function listIssues(
  projectId?: string,
  page = 1,
  limit = 20
): Promise<IssueListResponse> {
  const qs = new URLSearchParams({ page: String(page), limit: String(limit) });
  if (projectId) qs.set("project_id", projectId);
  return apiFetch<IssueListResponse>(`/api/v1/issues?${qs}`);
}

export async function getIssue(id: string): Promise<Issue> {
  return apiFetch<Issue>(`/api/v1/issues/${id}`);
}

export async function createIssue(payload: CreateIssuePayload): Promise<Issue> {
  return apiFetch<Issue>("/api/v1/issues", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function updateIssue(
  id: string,
  payload: UpdateIssuePayload
): Promise<Issue> {
  return apiFetch<Issue>(`/api/v1/issues/${id}`, {
    method: "PATCH",
    body: JSON.stringify(payload),
  });
}

export async function deleteIssue(id: string): Promise<void> {
  await apiFetch(`/api/v1/issues/${id}`, { method: "DELETE" });
}
