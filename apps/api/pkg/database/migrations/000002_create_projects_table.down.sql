-- Rollback: Projects/Organizations schema

DROP TRIGGER IF EXISTS update_projects_updated_at ON projects;
DROP TRIGGER IF EXISTS update_organizations_updated_at ON organizations;

DROP INDEX IF EXISTS idx_projects_key;
DROP INDEX IF EXISTS idx_projects_owner_id;
DROP INDEX IF EXISTS idx_projects_org_id;

DROP INDEX IF EXISTS idx_org_members_user_id;
DROP INDEX IF EXISTS idx_org_members_org_id;

DROP INDEX IF EXISTS idx_organizations_is_deleted;
DROP INDEX IF EXISTS idx_organizations_slug;
DROP INDEX IF EXISTS idx_organizations_owner_id;

DROP TABLE IF EXISTS team_members;
DROP TABLE IF EXISTS teams;
DROP TABLE IF EXISTS organization_members;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS organizations;
