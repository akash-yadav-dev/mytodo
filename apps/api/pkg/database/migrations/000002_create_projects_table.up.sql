-- Migration 000002: Organizations, Teams & Projects
-- Depends on: 000000_bootstrap_auth_tables (users table + update_updated_at_column fn)

CREATE TABLE IF NOT EXISTS organizations (
    id          UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    name        VARCHAR(255) NOT NULL,
    slug        VARCHAR(100) UNIQUE,
    description TEXT,
    plan_id     VARCHAR(50)  DEFAULT 'free',
    owner_id    UUID         NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    is_active   BOOLEAN      NOT NULL DEFAULT true,
    is_deleted  BOOLEAN      NOT NULL DEFAULT false,
    created_by  UUID         REFERENCES users(id),
    updated_by  UUID         REFERENCES users(id),
    deleted_by  UUID         REFERENCES users(id),
    created_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP    NOT NULL DEFAULT NOW(),
    deleted_at  TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_organizations_owner_id  ON organizations(owner_id);
CREATE INDEX IF NOT EXISTS idx_organizations_slug      ON organizations(slug);
CREATE INDEX IF NOT EXISTS idx_organizations_is_deleted ON organizations(is_deleted);

CREATE TABLE IF NOT EXISTS organization_members (
    id              UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID      NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id         UUID      NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role            VARCHAR(50) NOT NULL DEFAULT 'member',
    joined_at       TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP NOT NULL DEFAULT NOW(),
    invited_by      UUID      REFERENCES users(id),
    UNIQUE(organization_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_org_members_org_id  ON organization_members(organization_id);
CREATE INDEX IF NOT EXISTS idx_org_members_user_id ON organization_members(user_id);

CREATE TABLE IF NOT EXISTS teams (
    id              UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID         NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    slug            VARCHAR(100),
    description     TEXT,
    is_active       BOOLEAN      NOT NULL DEFAULT true,
    created_at      TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_teams_org_id ON teams(organization_id);

CREATE TABLE IF NOT EXISTS team_members (
    id        UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    team_id   UUID        NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id   UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role      VARCHAR(50) NOT NULL DEFAULT 'member',
    joined_at TIMESTAMP   NOT NULL DEFAULT NOW(),
    UNIQUE(team_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_team_members_team_id ON team_members(team_id);
CREATE INDEX IF NOT EXISTS idx_team_members_user_id ON team_members(user_id);

CREATE TABLE IF NOT EXISTS projects (
    id              UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID        REFERENCES organizations(id) ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    key             VARCHAR(10)  UNIQUE NOT NULL,
    description     TEXT,
    status          VARCHAR(50)  DEFAULT 'active',
    owner_id        UUID        REFERENCES users(id),
    created_at      TIMESTAMP   NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_projects_org_id   ON projects(organization_id);
CREATE INDEX IF NOT EXISTS idx_projects_owner_id ON projects(owner_id);
CREATE INDEX IF NOT EXISTS idx_projects_key      ON projects(key);

CREATE OR REPLACE TRIGGER update_organizations_updated_at
    BEFORE UPDATE ON organizations
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE OR REPLACE TRIGGER update_organization_members_updated_at
    BEFORE UPDATE ON organization_members
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE OR REPLACE TRIGGER update_teams_updated_at
    BEFORE UPDATE ON teams
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE OR REPLACE TRIGGER update_projects_updated_at
    BEFORE UPDATE ON projects
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Comments
COMMENT ON TABLE organizations        IS 'Top-level tenant organizations';
COMMENT ON TABLE organization_members IS 'Org membership with role-based access';
COMMENT ON TABLE teams                IS 'Sub-groups within an organization';
COMMENT ON TABLE projects             IS 'Project workspaces scoped to an organization';
