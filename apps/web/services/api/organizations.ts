import { apiFetch } from "@/services/api/client";
import type {
  Organization,
  OrganizationListResponse,
  CreateOrganizationPayload,
  UpdateOrganizationPayload,
  AddMemberPayload,
  Member,
} from "@/lib/types/organizations";

// ── Organizations ──────────────────────────────────────────────────────────

export async function listOrganizations(): Promise<OrganizationListResponse> {
  return apiFetch<OrganizationListResponse>("/api/v1/organizations");
}

export async function getOrganization(id: string): Promise<Organization> {
  return apiFetch<Organization>(`/api/v1/organizations/${id}`);
}

export async function createOrganization(payload: CreateOrganizationPayload): Promise<Organization> {
  return apiFetch<Organization>("/api/v1/organizations", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function updateOrganization(id: string, payload: UpdateOrganizationPayload): Promise<Organization> {
  return apiFetch<Organization>(`/api/v1/organizations/${id}`, {
    method: "PUT",
    body: JSON.stringify(payload),
  });
}

export async function deleteOrganization(id: string): Promise<void> {
  await apiFetch(`/api/v1/organizations/${id}`, { method: "DELETE" });
}

export async function getMyOrganizations(): Promise<OrganizationListResponse> {
  return apiFetch<OrganizationListResponse>("/api/v1/organizations/me/owned");
}

export async function getMemberOrganizations(): Promise<OrganizationListResponse> {
  return apiFetch<OrganizationListResponse>("/api/v1/organizations/me/member");
}

export async function transferOwnership(id: string, newOwnerId: string): Promise<Organization> {
  return apiFetch<Organization>(`/api/v1/organizations/${id}/transfer`, {
    method: "POST",
    body: JSON.stringify({ new_owner_id: newOwnerId }),
  });
}

// ── Members ────────────────────────────────────────────────────────────────

export async function listMembers(orgId: string): Promise<Member[]> {
  return apiFetch<Member[]>(`/api/v1/organizations/${orgId}/members`);
}

export async function addMember(orgId: string, payload: AddMemberPayload): Promise<Member> {
  return apiFetch<Member>(`/api/v1/organizations/${orgId}/members`, {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function removeMember(orgId: string, memberId: string): Promise<void> {
  await apiFetch(`/api/v1/organizations/${orgId}/members/${memberId}`, {
    method: "DELETE",
  });
}
