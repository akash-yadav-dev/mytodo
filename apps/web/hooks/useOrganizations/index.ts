import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import {
  listOrganizations,
  getOrganization,
  createOrganization,
  updateOrganization,
  deleteOrganization,
  getMyOrganizations,
  getMemberOrganizations,
  listMembers,
  addMember,
  removeMember,
} from "@/services/api/organizations";
import type {
  CreateOrganizationPayload,
  UpdateOrganizationPayload,
  AddMemberPayload,
} from "@/lib/types/organizations";

export const orgKeys = {
  all: ["organizations"] as const,
  lists: () => [...orgKeys.all, "list"] as const,
  detail: (id: string) => [...orgKeys.all, "detail", id] as const,
  members: (id: string) => [...orgKeys.all, "members", id] as const,
  mine: () => [...orgKeys.all, "mine"] as const,
  memberOf: () => [...orgKeys.all, "member-of"] as const,
};

export function useOrganizations() {
  return useQuery({ queryKey: orgKeys.lists(), queryFn: listOrganizations });
}

export function useOrganization(id: string) {
  return useQuery({
    queryKey: orgKeys.detail(id),
    queryFn: () => getOrganization(id),
    enabled: !!id,
  });
}

export function useOrganizationMembers(id: string) {
  return useQuery({
    queryKey: orgKeys.members(id),
    queryFn: () => listMembers(id),
    enabled: !!id,
  });
}

export function useMyOrganizations() {
  return useQuery({ queryKey: orgKeys.mine(), queryFn: getMyOrganizations });
}

export function useMemberOrganizations() {
  return useQuery({ queryKey: orgKeys.memberOf(), queryFn: getMemberOrganizations });
}

export function useCreateOrganization() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: CreateOrganizationPayload) => createOrganization(payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: orgKeys.lists() }),
  });
}

export function useUpdateOrganization(id: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: UpdateOrganizationPayload) => updateOrganization(id, payload),
    onSuccess: () => {
      qc.invalidateQueries({ queryKey: orgKeys.detail(id) });
      qc.invalidateQueries({ queryKey: orgKeys.lists() });
    },
  });
}

export function useDeleteOrganization() {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => deleteOrganization(id),
    onSuccess: () => qc.invalidateQueries({ queryKey: orgKeys.lists() }),
  });
}

export function useAddMember(orgId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (payload: AddMemberPayload) => addMember(orgId, payload),
    onSuccess: () => qc.invalidateQueries({ queryKey: orgKeys.members(orgId) }),
  });
}

export function useRemoveMember(orgId: string) {
  const qc = useQueryClient();
  return useMutation({
    mutationFn: (memberId: string) => removeMember(orgId, memberId),
    onSuccess: () => qc.invalidateQueries({ queryKey: orgKeys.members(orgId) }),
  });
}
