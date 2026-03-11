import { apiFetch } from "@/services/api/client";

export type Project = {
  id: string;
  name: string;
  key: string;
  description?: string;
  owner_id: string;
  org_id?: string;
  status?: "active" | "archived";
  created_at: string;
  updated_at: string;
};

export type ProjectListResponse = {
  data: Project[];
  total: number;
};

export type CreateProjectPayload = {
  name: string;
  key: string;
  description?: string;
  org_id?: string;
};

export type UpdateProjectPayload = Partial<CreateProjectPayload>;

export async function listProjects(): Promise<ProjectListResponse> {
  return apiFetch<ProjectListResponse>("/api/v1/projects");
}

export async function getProject(id: string): Promise<Project> {
  return apiFetch<Project>(`/api/v1/projects/${id}`);
}

export async function createProject(payload: CreateProjectPayload): Promise<Project> {
  return apiFetch<Project>("/api/v1/projects", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function updateProject(id: string, payload: UpdateProjectPayload): Promise<Project> {
  return apiFetch<Project>(`/api/v1/projects/${id}`, {
    method: "PUT",
    body: JSON.stringify(payload),
  });
}

export async function deleteProject(id: string): Promise<void> {
  await apiFetch(`/api/v1/projects/${id}`, { method: "DELETE" });
}
