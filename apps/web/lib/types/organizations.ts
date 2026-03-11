export type Organization = {
  id: string;
  name: string;
  slug?: string;
  description?: string;
  plan?: string;
  owner_id: string;
  is_active?: boolean;
  created_at: string;
  updated_at: string;
};

export type OrganizationListResponse = {
  data: Organization[];
  total?: number;
  page?: number;
  limit?: number;
};

export type CreateOrganizationPayload = {
  name: string;
  slug?: string;
  description?: string;
};

export type UpdateOrganizationPayload = Partial<CreateOrganizationPayload>;

export type TransferOwnershipPayload = {
  new_owner_id: string;
};

export type Member = {
  id: string;
  email: string;
  name?: string;
  role: string;
  joined_at?: string;
};

export type AddMemberPayload = {
  email: string;
  role?: string;
};
