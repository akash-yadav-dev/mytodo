# MyTodo Database Schema

**Version:** 1.0  
**Last Updated:** March 6, 2026  
**Database:** PostgreSQL 15+

This document describes the complete database schema for the MyTodo task management system, following best practices for a production-grade Jira-like application.

---

## 📋 Table of Contents

1. [Schema Overview](#schema-overview)
2. [Core Tables](#core-tables)
3. [User & Authentication](#user--authentication)
4. [Organizations & Projects](#organizations--projects)
5. [Issues & Tasks](#issues--tasks)
6. [Boards & Sprints](#boards--sprints)
7. [Comments & Attachments](#comments--attachments)
8. [Workflows & Automation](#workflows--automation)
9. [Notifications & Integrations](#notifications--integrations)
10. [Indexing Strategy](#indexing-strategy)
11. [Performance Optimizations](#performance-optimizations)

---

## Schema Overview

### Design Principles

- **Normalization**: 3NF for data integrity
- **Soft Deletes**: Preserve historical data
- **Audit Trails**: Track all changes
- **Partitioning**: For high-volume tables (issues, activities)
- **UUID Primary Keys**: Distributed system compatibility
- **JSONB**: Flexible metadata storage
- **Full-Text Search**: PostgreSQL tsvector for search

### Database Statistics (Target Scale)

| Metric | Target |
|--------|--------|
| Total Users | 5M+ |
| Organizations | 100K+ |
| Projects | 500K+ |
| Issues | 1B+ |
| Daily Transactions | 100M+ |

---

## Core Tables

### users

Stores user account information with authentication details.

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255), -- NULL if OAuth only
    full_name VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    timezone VARCHAR(50) DEFAULT 'UTC',
    locale VARCHAR(10) DEFAULT 'en-US',
    email_verified BOOLEAN DEFAULT FALSE,
    email_verified_at TIMESTAMP,
    phone_number VARCHAR(20),
    phone_verified BOOLEAN DEFAULT FALSE,
    two_factor_enabled BOOLEAN DEFAULT FALSE,
    two_factor_secret VARCHAR(255),
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'suspended', 'deleted')),
    last_login_at TIMESTAMP,
    last_login_ip INET,
    login_count INTEGER DEFAULT 0,
    failed_login_attempts INTEGER DEFAULT 0,
    locked_until TIMESTAMP,
    preferences JSONB DEFAULT '{}', -- UI preferences, notifications settings
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Indexes
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_username ON users(username) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_status ON users(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_last_login ON users(last_login_at);

-- Full-text search
CREATE INDEX idx_users_search ON users USING gin(to_tsvector('english', full_name || ' ' || email || ' ' || username));

COMMENT ON TABLE users IS 'User accounts with authentication and profile data';
```

**Improvements:**
- Added `status` enum for account lifecycle
- Added `locked_until` for security lockouts
- Added `login_count` and `last_login_ip` for analytics
- Added `preferences` JSONB for flexible user settings
- Added full-text search capability

---

### oauth_providers

Stores OAuth provider connections.

```sql
CREATE TABLE oauth_providers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    provider VARCHAR(50) NOT NULL, -- 'google', 'github', 'microsoft', etc.
    provider_user_id VARCHAR(255) NOT NULL,
    access_token TEXT,
    refresh_token TEXT,
    token_expires_at TIMESTAMP,
    scope TEXT,
    profile_data JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider, provider_user_id)
);

CREATE INDEX idx_oauth_user ON oauth_providers(user_id);
CREATE INDEX idx_oauth_provider ON oauth_providers(provider, provider_user_id);

COMMENT ON TABLE oauth_providers IS 'OAuth provider connections for social login';
```

---

### sessions

User authentication sessions.

```sql
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    token_hash VARCHAR(255) UNIQUE NOT NULL,
    refresh_token_hash VARCHAR(255),
    user_agent TEXT,
    ip_address INET,
    device_info JSONB DEFAULT '{}',
    expires_at TIMESTAMP NOT NULL,
    last_activity_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP
);

CREATE INDEX idx_sessions_user ON sessions(user_id);
CREATE INDEX idx_sessions_token ON sessions(token_hash) WHERE revoked_at IS NULL;
CREATE INDEX idx_sessions_expires ON sessions(expires_at) WHERE revoked_at IS NULL;
CREATE INDEX idx_sessions_activity ON sessions(last_activity_at);

COMMENT ON TABLE sessions IS 'Active user sessions with JWT tokens';
```

**Improvements:**
- Added `device_info` for multi-device support
- Added `last_activity_at` for session management
- Added `revoked_at` for manual session termination

---

## Organizations & Projects

### organizations

Multi-tenant organization accounts.

```sql
CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    logo_url TEXT,
    website_url TEXT,
    billing_email VARCHAR(255),
    plan_type VARCHAR(50) DEFAULT 'free' CHECK (plan_type IN ('free', 'starter', 'professional', 'enterprise')),
    plan_expires_at TIMESTAMP,
    max_users INTEGER DEFAULT 10,
    max_projects INTEGER DEFAULT 3,
    storage_limit_bytes BIGINT DEFAULT 1073741824, -- 1GB
    storage_used_bytes BIGINT DEFAULT 0,
    settings JSONB DEFAULT '{}',
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_orgs_slug ON organizations(slug) WHERE deleted_at IS NULL;
CREATE INDEX idx_orgs_plan ON organizations(plan_type);
CREATE INDEX idx_orgs_created ON organizations(created_at);

COMMENT ON TABLE organizations IS 'Organizations for multi-tenancy';
```

**Improvements:**
- Added `plan_type` and billing limits
- Added storage tracking
- Added `settings` JSONB for org-level configuration

---

### organization_members

Organization membership and roles.

```sql
CREATE TABLE organization_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL CHECK (role IN ('owner', 'admin', 'member', 'guest')),
    permissions JSONB DEFAULT '[]', -- Custom permissions array
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'inactive', 'pending')),
    invited_by UUID REFERENCES users(id),
    invitation_token VARCHAR(255),
    invitation_expires_at TIMESTAMP,
    joined_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(organization_id, user_id)
);

CREATE INDEX idx_org_members_org ON organization_members(organization_id);
CREATE INDEX idx_org_members_user ON organization_members(user_id);
CREATE INDEX idx_org_members_role ON organization_members(organization_id, role);
CREATE INDEX idx_org_members_invitation ON organization_members(invitation_token) WHERE invitation_expires_at > CURRENT_TIMESTAMP;

COMMENT ON TABLE organization_members IS 'Organization membership with roles';
```

**Improvements:**
- Added invitation workflow support
- Added `permissions` JSONB for granular access control
- Added `status` for pending invitations

---

### projects

Project workspaces within organizations.

```sql
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    key VARCHAR(10) NOT NULL, -- e.g., 'PROJ', 'TASK'
    description TEXT,
    icon VARCHAR(50),
    color VARCHAR(7), -- Hex color code
    visibility VARCHAR(20) DEFAULT 'private' CHECK (visibility IN ('public', 'private', 'internal')),
    default_assignee_id UUID REFERENCES users(id),
    lead_id UUID REFERENCES users(id),
    category VARCHAR(50), -- 'software', 'marketing', 'design', etc.
    start_date DATE,
    end_date DATE,
    status VARCHAR(20) DEFAULT 'active' CHECK (status IN ('active', 'archived', 'on_hold', 'cancelled')),
    settings JSONB DEFAULT '{}', -- Project-specific settings
    metadata JSONB DEFAULT '{}',
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    archived_at TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(organization_id, key)
);

CREATE INDEX idx_projects_org ON projects(organization_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_projects_key ON projects(organization_id, key);
CREATE INDEX idx_projects_lead ON projects(lead_id);
CREATE INDEX idx_projects_status ON projects(status) WHERE deleted_at IS NULL;
CREATE INDEX idx_projects_search ON projects USING gin(to_tsvector('english', name || ' ' || COALESCE(description, '')));

COMMENT ON TABLE projects IS 'Project workspaces for organizing issues';
```

**Improvements:**
- Added `category` for project classification
- Added `visibility` for access control
- Added date tracking for project timeline
- Added full-text search

---

### project_members

Project team members with roles.

```sql
CREATE TABLE project_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL CHECK (role IN ('lead', 'admin', 'developer', 'reporter', 'viewer')),
    permissions JSONB DEFAULT '[]',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, user_id)
);

CREATE INDEX idx_project_members_project ON project_members(project_id);
CREATE INDEX idx_project_members_user ON project_members(user_id);
CREATE INDEX idx_project_members_role ON project_members(project_id, role);

COMMENT ON TABLE project_members IS 'Project team membership';
```

---

## Issues & Tasks

### issue_types

Configurable issue types (Story, Bug, Task, Epic, etc.).

```sql
CREATE TABLE issue_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE, -- NULL for org-wide types
    name VARCHAR(50) NOT NULL,
    description TEXT,
    icon VARCHAR(50),
    color VARCHAR(7),
    is_subtask BOOLEAN DEFAULT FALSE,
    is_system BOOLEAN DEFAULT FALSE, -- System-defined types
    position INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_issue_types_org ON issue_types(organization_id);
CREATE INDEX idx_issue_types_project ON issue_types(project_id);
CREATE INDEX idx_issue_types_position ON issue_types(organization_id, position);

COMMENT ON TABLE issue_types IS 'Configurable issue types';

-- Default issue types
INSERT INTO issue_types (id, name, icon, color, is_system) VALUES
    (gen_random_uuid(), 'Epic', '📚', '#8B5CF6', TRUE),
    (gen_random_uuid(), 'Story', '📖', '#10B981', TRUE),
    (gen_random_uuid(), 'Task', '✓', '#3B82F6', TRUE),
    (gen_random_uuid(), 'Bug', '🐛', '#EF4444', TRUE),
    (gen_random_uuid(), 'Subtask', '📝', '#6B7280', TRUE);
```

---

### issue_statuses

Workflow statuses.

```sql
CREATE TABLE issue_statuses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    category VARCHAR(20) NOT NULL CHECK (category IN ('todo', 'in_progress', 'done', 'cancelled')),
    color VARCHAR(7),
    position INTEGER DEFAULT 0,
    is_default BOOLEAN DEFAULT FALSE,
    is_system BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_statuses_org ON issue_statuses(organization_id);
CREATE INDEX idx_statuses_project ON issue_statuses(project_id);
CREATE INDEX idx_statuses_category ON issue_statuses(category);
CREATE INDEX idx_statuses_position ON issue_statuses(project_id, position);

COMMENT ON TABLE issue_statuses IS 'Issue workflow statuses';

-- Default statuses
INSERT INTO issue_statuses (id, name, category, color, is_system) VALUES
    (gen_random_uuid(), 'Backlog', 'todo', '#6B7280', TRUE),
    (gen_random_uuid(), 'To Do', 'todo', '#3B82F6', TRUE),
    (gen_random_uuid(), 'In Progress', 'in_progress', '#F59E0B', TRUE),
    (gen_random_uuid(), 'In Review', 'in_progress', '#8B5CF6', TRUE),
    (gen_random_uuid(), 'Done', 'done', '#10B981', TRUE),
    (gen_random_uuid(), 'Cancelled', 'cancelled', '#EF4444', TRUE);
```

---

### issue_priorities

Issue priority levels.

```sql
CREATE TABLE issue_priorities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID REFERENCES organizations(id) ON DELETE CASCADE,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    level INTEGER NOT NULL, -- 1 = Highest, 5 = Lowest
    icon VARCHAR(50),
    color VARCHAR(7),
    is_system BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_priorities_org ON issue_priorities(organization_id);
CREATE INDEX idx_priorities_level ON issue_priorities(level);

COMMENT ON TABLE issue_priorities IS 'Issue priority levels';

-- Default priorities
INSERT INTO issue_priorities (id, name, level, icon, color, is_system) VALUES
    (gen_random_uuid(), 'Critical', 1, '🔴', '#DC2626', TRUE),
    (gen_random_uuid(), 'High', 2, '🟠', '#F59E0B', TRUE),
    (gen_random_uuid(), 'Medium', 3, '🟡', '#FCD34D', TRUE),
    (gen_random_uuid(), 'Low', 4, '🟢', '#10B981', TRUE),
    (gen_random_uuid(), 'Trivial', 5, '⚪', '#9CA3AF', TRUE);
```

---

### issues

**Main issues table - partition by created_at for scalability.**

```sql
CREATE TABLE issues (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    issue_number SERIAL NOT NULL, -- Auto-increment per project
    issue_key VARCHAR(50) NOT NULL, -- e.g., 'PROJ-123'
    title VARCHAR(500) NOT NULL,
    description TEXT,
    description_html TEXT,
    description_search TSVECTOR, -- Full-text search vector
    
    -- Classification
    issue_type_id UUID NOT NULL REFERENCES issue_types(id),
    status_id UUID NOT NULL REFERENCES issue_statuses(id),
    priority_id UUID REFERENCES issue_priorities(id),
    
    -- Assignment
    reporter_id UUID NOT NULL REFERENCES users(id),
    assignee_id UUID REFERENCES users(id),
    
    -- Hierarchy
    parent_issue_id UUID REFERENCES issues(id),
    epic_id UUID REFERENCES issues(id),
    
    -- Sprint & Board
    sprint_id UUID REFERENCES sprints(id),
    board_column_id UUID,
    board_position INTEGER,
    
    -- Estimation
    story_points INTEGER CHECK (story_points >= 0),
    estimate_minutes INTEGER CHECK (estimate_minutes >= 0),
    time_spent_minutes INTEGER DEFAULT 0,
    remaining_estimate_minutes INTEGER,
    
    -- Dates
    due_date TIMESTAMP,
    start_date TIMESTAMP,
    resolved_at TIMESTAMP,
    closed_at TIMESTAMP,
    
    -- Additional fields
    environment TEXT, -- For bugs
    security_level VARCHAR(20) DEFAULT 'public',
    
    -- Counters (denormalized for performance)
    comment_count INTEGER DEFAULT 0,
    attachment_count INTEGER DEFAULT 0,
    subtask_count INTEGER DEFAULT 0,
    watcher_count INTEGER DEFAULT 0,
    vote_count INTEGER DEFAULT 0,
    
    -- Metadata
    labels TEXT[], -- Array of label names for quick filtering
    custom_fields JSONB DEFAULT '{}',
    metadata JSONB DEFAULT '{}',
    
    -- Audit
    created_by UUID REFERENCES users(id),
    updated_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    
    UNIQUE(project_id, issue_number)
) PARTITION BY RANGE (created_at);

-- Create partitions (example for 2026)
CREATE TABLE issues_2026_q1 PARTITION OF issues
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');
CREATE TABLE issues_2026_q2 PARTITION OF issues
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');
CREATE TABLE issues_2026_q3 PARTITION OF issues
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');
CREATE TABLE issues_2026_q4 PARTITION OF issues
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

-- Indexes
CREATE INDEX idx_issues_project ON issues(project_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_org ON issues(organization_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_key ON issues(issue_key) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_assignee ON issues(assignee_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_reporter ON issues(reporter_id);
CREATE INDEX idx_issues_status ON issues(status_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_priority ON issues(priority_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_sprint ON issues(sprint_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_epic ON issues(epic_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_parent ON issues(parent_issue_id);
CREATE INDEX idx_issues_due_date ON issues(due_date) WHERE deleted_at IS NULL AND due_date IS NOT NULL;
CREATE INDEX idx_issues_created ON issues(created_at);
CREATE INDEX idx_issues_updated ON issues(updated_at) WHERE deleted_at IS NULL;

-- GIN indexes for arrays and JSONB
CREATE INDEX idx_issues_labels ON issues USING gin(labels);
CREATE INDEX idx_issues_custom_fields ON issues USING gin(custom_fields);

-- Full-text search
CREATE INDEX idx_issues_search ON issues USING gin(description_search);

-- Composite indexes for common queries
CREATE INDEX idx_issues_project_status ON issues(project_id, status_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_project_assignee ON issues(project_id, assignee_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_issues_sprint_status ON issues(sprint_id, status_id) WHERE deleted_at IS NULL AND sprint_id IS NOT NULL;

-- Trigger to update search vector
CREATE OR REPLACE FUNCTION update_issue_search_vector()
RETURNS TRIGGER AS $$
BEGIN
    NEW.description_search := to_tsvector('english', 
        COALESCE(NEW.title, '') || ' ' || 
        COALESCE(NEW.description, '') || ' ' || 
        COALESCE(NEW.issue_key, '')
    );
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_issue_search
    BEFORE INSERT OR UPDATE ON issues
    FOR EACH ROW
    EXECUTE FUNCTION update_issue_search_vector();

COMMENT ON TABLE issues IS 'Main issues/tasks table with partitioning for scalability';
```

**Major Improvements:**
- **Partitioning**: By created_at for handling billions of issues
- **Denormalized counters**: For performance (comment_count, etc.)
- **Full-text search**: With tsvector and trigger
- **Comprehensive indexing**: For all common query patterns
- **Labels array**: For quick filtering without joins
- **Custom fields**: JSONB for flexibility
- **Board position**: For drag-drop ordering

---

### custom_fields

Custom field definitions.

```sql
CREATE TABLE custom_fields (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    field_type VARCHAR(50) NOT NULL CHECK (field_type IN (
        'text', 'number', 'date', 'datetime', 'boolean', 
        'select', 'multi_select', 'user', 'url', 'email'
    )),
    options JSONB DEFAULT '[]', -- For select/multi_select
    default_value TEXT,
    is_required BOOLEAN DEFAULT FALSE,
    validation_rules JSONB DEFAULT '{}',
    position INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_custom_fields_org ON custom_fields(organization_id);
CREATE INDEX idx_custom_fields_project ON custom_fields(project_id);

COMMENT ON TABLE custom_fields IS 'Custom field definitions';
```

---

### issue_links

Relationships between issues.

```sql
CREATE TABLE issue_links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    source_issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    target_issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    link_type VARCHAR(50) NOT NULL CHECK (link_type IN (
        'blocks', 'is_blocked_by',
        'relates_to',
        'duplicates', 'is_duplicated_by',
        'causes', 'is_caused_by',
        'parent_of', 'child_of'
    )),
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(source_issue_id, target_issue_id, link_type)
);

CREATE INDEX idx_issue_links_source ON issue_links(source_issue_id);
CREATE INDEX idx_issue_links_target ON issue_links(target_issue_id);
CREATE INDEX idx_issue_links_type ON issue_links(link_type);

COMMENT ON TABLE issue_links IS 'Issue relationships and dependencies';
```

---

### issue_history

Complete audit trail of all issue changes.

```sql
CREATE TABLE issue_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    field_name VARCHAR(100) NOT NULL,
    field_type VARCHAR(50),
    old_value TEXT,
    new_value TEXT,
    old_value_data JSONB,
    new_value_data JSONB,
    changed_by UUID REFERENCES users(id),
    changed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) PARTITION BY RANGE (changed_at);

-- Create partitions
CREATE TABLE issue_history_2026_q1 PARTITION OF issue_history
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');
CREATE TABLE issue_history_2026_q2 PARTITION OF issue_history
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');
CREATE TABLE issue_history_2026_q3 PARTITION OF issue_history
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');
CREATE TABLE issue_history_2026_q4 PARTITION OF issue_history
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

CREATE INDEX idx_history_issue ON issue_history(issue_id);
CREATE INDEX idx_history_field ON issue_history(field_name);
CREATE INDEX idx_history_timestamp ON issue_history(changed_at);
CREATE INDEX idx_history_user ON issue_history(changed_by);

COMMENT ON TABLE issue_history IS 'Complete audit trail of issue changes';
```

---

### labels

Reusable labels for categorization.

```sql
CREATE TABLE labels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    color VARCHAR(7) NOT NULL DEFAULT '#6B7280',
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(organization_id, COALESCE(project_id::text, 'global'), name)
);

CREATE INDEX idx_labels_org ON labels(organization_id);
CREATE INDEX idx_labels_project ON labels(project_id);
CREATE INDEX idx_labels_name ON labels(organization_id, name);

COMMENT ON TABLE labels IS 'Reusable labels for issue categorization';
```

---

### issue_labels

Many-to-many relationship between issues and labels.

```sql
CREATE TABLE issue_labels (
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    label_id UUID NOT NULL REFERENCES labels(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (issue_id, label_id)
);

CREATE INDEX idx_issue_labels_label ON issue_labels(label_id);

COMMENT ON TABLE issue_labels IS 'Issue label assignments';
```

---

### issue_watchers

Users watching issues for notifications.

```sql
CREATE TABLE issue_watchers (
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (issue_id, user_id)
);

CREATE INDEX idx_watchers_user ON issue_watchers(user_id);

COMMENT ON TABLE issue_watchers IS 'Users watching issues';
```

---

### issue_votes

Issue upvoting/voting.

```sql
CREATE TABLE issue_votes (
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    vote INTEGER CHECK (vote IN (-1, 1)), -- -1 = downvote, 1 = upvote
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (issue_id, user_id)
);

CREATE INDEX idx_votes_user ON issue_votes(user_id);

COMMENT ON TABLE issue_votes IS 'Issue voting';
```

---

## Boards & Sprints

### boards

Kanban/Scrum boards.

```sql
CREATE TABLE boards (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    board_type VARCHAR(20) NOT NULL CHECK (board_type IN ('scrum', 'kanban')),
    is_default BOOLEAN DEFAULT FALSE,
    filter_config JSONB DEFAULT '{}', -- Saved filters
    settings JSONB DEFAULT '{}',
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_boards_project ON boards(project_id);
CREATE INDEX idx_boards_type ON boards(board_type);

COMMENT ON TABLE boards IS 'Kanban and Scrum boards';
```

---

### board_columns

Board columns with positions.

```sql
CREATE TABLE board_columns (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    board_id UUID NOT NULL REFERENCES boards(id) ON DELETE CASCADE,
    status_id UUID NOT NULL REFERENCES issue_statuses(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    position INTEGER NOT NULL,
    wip_limit INTEGER, -- Work In Progress limit
    settings JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_board_columns_board ON board_columns(board_id);
CREATE INDEX idx_board_columns_status ON board_columns(status_id);
CREATE INDEX idx_board_columns_position ON board_columns(board_id, position);

COMMENT ON TABLE board_columns IS 'Board columns with status mappings';
```

---

### sprints

Agile sprints for time-boxed iterations.

```sql
CREATE TABLE sprints (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    board_id UUID REFERENCES boards(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    goal TEXT,
    sprint_number INTEGER,
    status VARCHAR(20) DEFAULT 'planned' CHECK (status IN ('planned', 'active', 'completed', 'cancelled')),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    completed_at TIMESTAMP,
    
    -- Sprint metrics (denormalized)
    total_points INTEGER DEFAULT 0,
    completed_points INTEGER DEFAULT 0,
    total_issues INTEGER DEFAULT 0,
    completed_issues INTEGER DEFAULT 0,
    
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_sprints_project ON sprints(project_id);
CREATE INDEX idx_sprints_board ON sprints(board_id);
CREATE INDEX idx_sprints_status ON sprints(status);
CREATE INDEX idx_sprints_dates ON sprints(start_date, end_date);

COMMENT ON TABLE sprints IS 'Agile sprints for iterative development';
```

---

## Comments & Attachments

### comments

Issue comments with threading.

```sql
CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    parent_comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    content_html TEXT,
    is_internal BOOLEAN DEFAULT FALSE, -- Internal team notes
    edited BOOLEAN DEFAULT FALSE,
    edited_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
) PARTITION BY RANGE (created_at);

-- Create partitions
CREATE TABLE comments_2026_q1 PARTITION OF comments
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');
CREATE TABLE comments_2026_q2 PARTITION OF comments
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');
CREATE TABLE comments_2026_q3 PARTITION OF comments
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');
CREATE TABLE comments_2026_q4 PARTITION OF comments
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

CREATE INDEX idx_comments_issue ON comments(issue_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_comments_user ON comments(user_id);
CREATE INDEX idx_comments_parent ON comments(parent_comment_id);
CREATE INDEX idx_comments_created ON comments(created_at);

COMMENT ON TABLE comments IS 'Issue comments with threading support';
```

---

### attachments

File attachments for issues.

```sql
CREATE TABLE attachments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    issue_id UUID REFERENCES issues(id) ON DELETE CASCADE,
    comment_id UUID REFERENCES comments(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    file_type VARCHAR(100),
    storage_provider VARCHAR(50) DEFAULT 's3', -- 's3', 'minio', 'local'
    storage_path TEXT NOT NULL,
    storage_key TEXT UNIQUE NOT NULL,
    thumbnail_path TEXT,
    download_count INTEGER DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CHECK (issue_id IS NOT NULL OR comment_id IS NOT NULL)
);

CREATE INDEX idx_attachments_issue ON attachments(issue_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attachments_comment ON attachments(comment_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_attachments_user ON attachments(user_id);
CREATE INDEX idx_attachments_created ON attachments(created_at);

COMMENT ON TABLE attachments IS 'File attachments for issues and comments';
```

---

### worklogs

Time tracking entries.

```sql
CREATE TABLE worklogs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    issue_id UUID NOT NULL REFERENCES issues(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),
    time_spent_minutes INTEGER NOT NULL CHECK (time_spent_minutes > 0),
    comment TEXT,
    started_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_worklogs_issue ON worklogs(issue_id) WHERE deleted_at IS NULL;
CREATE INDEX idx_worklogs_user ON worklogs(user_id);
CREATE INDEX idx_worklogs_started ON worklogs(started_at);

COMMENT ON TABLE worklogs IS 'Time tracking entries for issues';
```

---

## Workflows & Automation

### workflows

Custom workflow definitions.

```sql
CREATE TABLE workflows (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    is_default BOOLEAN DEFAULT FALSE,
    is_active BOOLEAN DEFAULT TRUE,
    config JSONB DEFAULT '{}',
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_workflows_org ON workflows(organization_id);
CREATE INDEX idx_workflows_project ON workflows(project_id);

COMMENT ON TABLE workflows IS 'Custom workflow definitions';
```

---

### workflow_transitions

Allowed status transitions.

```sql
CREATE TABLE workflow_transitions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(id) ON DELETE CASCADE,
    from_status_id UUID NOT NULL REFERENCES issue_statuses(id),
    to_status_id UUID NOT NULL REFERENCES issue_statuses(id),
    name VARCHAR(100),
    required_fields TEXT[],
    validators JSONB DEFAULT '[]',
    post_functions JSONB DEFAULT '[]',
    conditions JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_transitions_workflow ON workflow_transitions(workflow_id);
CREATE INDEX idx_transitions_from ON workflow_transitions(from_status_id);
CREATE INDEX idx_transitions_to ON workflow_transitions(to_status_id);

COMMENT ON TABLE workflow_transitions IS 'Workflow status transitions';
```

---

### automation_rules

Automation rules for issues.

```sql
CREATE TABLE automation_rules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    trigger_type VARCHAR(50) NOT NULL CHECK (trigger_type IN (
        'issue_created', 'issue_updated', 'issue_transitioned',
        'comment_added', 'field_changed', 'scheduled', 'manual'
    )),
    trigger_config JSONB DEFAULT '{}',
    conditions JSONB DEFAULT '[]',
    actions JSONB DEFAULT '[]',
    is_active BOOLEAN DEFAULT TRUE,
    execution_count INTEGER DEFAULT 0,
    last_executed_at TIMESTAMP,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_automation_org ON automation_rules(organization_id);
CREATE INDEX idx_automation_project ON automation_rules(project_id);
CREATE INDEX idx_automation_trigger ON automation_rules(trigger_type) WHERE is_active = TRUE;

COMMENT ON TABLE automation_rules IS 'Issue automation rules';
```

---

## Notifications & Integrations

### notifications

User notifications.

```sql
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    message TEXT,
    link_url TEXT,
    entity_type VARCHAR(50), -- 'issue', 'comment', 'project', etc.
    entity_id UUID,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    priority VARCHAR(20) DEFAULT 'normal' CHECK (priority IN ('low', 'normal', 'high', 'urgent')),
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) PARTITION BY RANGE (created_at);

-- Create partitions
CREATE TABLE notifications_2026_q1 PARTITION OF notifications
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');
CREATE TABLE notifications_2026_q2 PARTITION OF notifications
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');
CREATE TABLE notifications_2026_q3 PARTITION OF notifications
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');
CREATE TABLE notifications_2026_q4 PARTITION OF notifications
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

CREATE INDEX idx_notifications_user ON notifications(user_id) WHERE is_read = FALSE;
CREATE INDEX idx_notifications_entity ON notifications(entity_type, entity_id);
CREATE INDEX idx_notifications_created ON notifications(created_at);

COMMENT ON TABLE notifications IS 'User notifications';
```

---

### integrations

Third-party integrations.

```sql
CREATE TABLE integrations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    integration_type VARCHAR(50) NOT NULL CHECK (integration_type IN (
        'github', 'gitlab', 'slack', 'discord', 'webhook',
        'jira', 'trello', 'email', 'calendar'
    )),
    name VARCHAR(255) NOT NULL,
    config JSONB NOT NULL DEFAULT '{}',
    credentials JSONB, -- Encrypted
    is_active BOOLEAN DEFAULT TRUE,
    last_sync_at TIMESTAMP,
    sync_status VARCHAR(20),
    error_message TEXT,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_integrations_org ON integrations(organization_id);
CREATE INDEX idx_integrations_project ON integrations(project_id);
CREATE INDEX idx_integrations_type ON integrations(integration_type) WHERE is_active = TRUE;

COMMENT ON TABLE integrations IS 'Third-party integrations';
```

---

### webhooks

Webhook endpoints.

```sql
CREATE TABLE webhooks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    secret VARCHAR(255),
    events TEXT[] NOT NULL, -- Array of event types
    is_active BOOLEAN DEFAULT TRUE,
    ssl_verify BOOLEAN DEFAULT TRUE,
    timeout_seconds INTEGER DEFAULT 30,
    retry_count INTEGER DEFAULT 3,
    last_triggered_at TIMESTAMP,
    failure_count INTEGER DEFAULT 0,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_webhooks_org ON webhooks(organization_id);
CREATE INDEX idx_webhooks_project ON webhooks(project_id);
CREATE INDEX idx_webhooks_active ON webhooks(is_active);

COMMENT ON TABLE webhooks IS 'Webhook endpoints for event notifications';
```

---

### webhook_deliveries

Webhook delivery logs.

```sql
CREATE TABLE webhook_deliveries (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    webhook_id UUID NOT NULL REFERENCES webhooks(id) ON DELETE CASCADE,
    event_type VARCHAR(100) NOT NULL,
    payload JSONB NOT NULL,
    response_code INTEGER,
    response_body TEXT,
    error_message TEXT,
    duration_ms INTEGER,
    attempt_number INTEGER DEFAULT 1,
    delivered BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
) PARTITION BY RANGE (created_at);

-- Create partitions
CREATE TABLE webhook_deliveries_2026_q1 PARTITION OF webhook_deliveries
    FOR VALUES FROM ('2026-01-01') TO ('2026-04-01');
CREATE TABLE webhook_deliveries_2026_q2 PARTITION OF webhook_deliveries
    FOR VALUES FROM ('2026-04-01') TO ('2026-07-01');
CREATE TABLE webhook_deliveries_2026_q3 PARTITION OF webhook_deliveries
    FOR VALUES FROM ('2026-07-01') TO ('2026-10-01');
CREATE TABLE webhook_deliveries_2026_q4 PARTITION OF webhook_deliveries
    FOR VALUES FROM ('2026-10-01') TO ('2027-01-01');

CREATE INDEX idx_webhook_deliveries_webhook ON webhook_deliveries(webhook_id);
CREATE INDEX idx_webhook_deliveries_event ON webhook_deliveries(event_type);
CREATE INDEX idx_webhook_deliveries_delivered ON webhook_deliveries(delivered);
CREATE INDEX idx_webhook_deliveries_created ON webhook_deliveries(created_at);

COMMENT ON TABLE webhook_deliveries IS 'Webhook delivery logs';
```

---

## Indexing Strategy

### Composite Indexes for Common Queries

```sql
-- Dashboard queries
CREATE INDEX idx_issues_user_dashboard ON issues(assignee_id, status_id, updated_at DESC) 
    WHERE deleted_at IS NULL;

-- Project queries
CREATE INDEX idx_issues_project_sprint ON issues(project_id, sprint_id, status_id) 
    WHERE deleted_at IS NULL AND sprint_id IS NOT NULL;

-- Search queries
CREATE INDEX idx_issues_org_search ON issues(organization_id) 
    INCLUDE (title, issue_key, status_id, assignee_id) 
    WHERE deleted_at IS NULL;
```

---

## Performance Optimizations

### 1. Materialized Views

```sql
-- Sprint velocity view
CREATE MATERIALIZED VIEW sprint_velocity AS
SELECT 
    s.id AS sprint_id,
    s.project_id,
    COUNT(i.id) AS total_issues,
    SUM(CASE WHEN i.status_id IN (SELECT id FROM issue_statuses WHERE category = 'done') THEN 1 ELSE 0 END) AS completed_issues,
    SUM(i.story_points) AS total_points,
    SUM(CASE WHEN i.status_id IN (SELECT id FROM issue_statuses WHERE category = 'done') THEN i.story_points ELSE 0 END) AS completed_points
FROM sprints s
LEFT JOIN issues i ON i.sprint_id = s.id
WHERE s.status != 'cancelled'
GROUP BY s.id, s.project_id;

CREATE UNIQUE INDEX idx_sprint_velocity_sprint ON sprint_velocity(sprint_id);

-- Refresh periodically
CREATE OR REPLACE FUNCTION refresh_sprint_velocity()
RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY sprint_velocity;
END;
$$ LANGUAGE plpgsql;
```

---

### 2. Database Functions

```sql
-- Function to get issue activity
CREATE OR REPLACE FUNCTION get_issue_activity(p_issue_id UUID, p_limit INTEGER DEFAULT 50)
RETURNS TABLE (
    activity_type VARCHAR,
    activity_data JSONB,
    user_id UUID,
    created_at TIMESTAMP
) AS $$
BEGIN
    RETURN QUERY
    -- Comments
    SELECT 
        'comment'::VARCHAR,
        jsonb_build_object('id', c.id, 'content', c.content),
        c.user_id,
        c.created_at
    FROM comments c
    WHERE c.issue_id = p_issue_id AND c.deleted_at IS NULL
    
    UNION ALL
    
    -- History
    SELECT 
        'field_change'::VARCHAR,
        jsonb_build_object('field', h.field_name, 'old', h.old_value, 'new', h.new_value),
        h.changed_by,
        h.changed_at
    FROM issue_history h
    WHERE h.issue_id = p_issue_id
    
    UNION ALL
    
    -- Worklogs
    SELECT 
        'worklog'::VARCHAR,
        jsonb_build_object('id', w.id, 'time_spent', w.time_spent_minutes),
        w.user_id,
        w.created_at
    FROM worklogs w
    WHERE w.issue_id = p_issue_id AND w.deleted_at IS NULL
    
    ORDER BY created_at DESC
    LIMIT p_limit;
END;
$$ LANGUAGE plpgsql;
```

---

### 3. Triggers for Denormalization

```sql
-- Update issue counters
CREATE OR REPLACE FUNCTION update_issue_counters()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_TABLE_NAME = 'comments' THEN
        UPDATE issues SET comment_count = comment_count + 1 WHERE id = NEW.issue_id;
    ELSIF TG_TABLE_NAME = 'attachments' THEN
        UPDATE issues SET attachment_count = attachment_count + 1 WHERE id = NEW.issue_id;
    ELSIF TG_TABLE_NAME = 'issue_watchers' THEN
        UPDATE issues SET watcher_count = watcher_count + 1 WHERE id = NEW.issue_id;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_comment_count
    AFTER INSERT ON comments
    FOR EACH ROW
    EXECUTE FUNCTION update_issue_counters();

CREATE TRIGGER trigger_update_attachment_count
    AFTER INSERT ON attachments
    FOR EACH ROW
    EXECUTE FUNCTION update_issue_counters();

CREATE TRIGGER trigger_update_watcher_count
    AFTER INSERT ON issue_watchers
    FOR EACH ROW
    EXECUTE FUNCTION update_issue_counters();
```

---

### 4. Partitioning Strategy

**Tables to Partition:**
- `issues` - By created_at (quarterly)
- `issue_history` - By changed_at (quarterly)
- `comments` - By created_at (quarterly)
- `notifications` - By created_at (monthly)
- `webhook_deliveries` - By created_at (monthly)

**Benefits:**
- Faster queries with partition pruning
- Easier archival and purging
- Better maintenance (VACUUM, ANALYZE)

---

### 5. Query Optimization Tips

```sql
-- Use EXPLAIN ANALYZE for slow queries
EXPLAIN ANALYZE
SELECT * FROM issues 
WHERE project_id = 'xxx' 
AND status_id = 'yyy' 
AND deleted_at IS NULL;

-- Create statistics for better query planning
CREATE STATISTICS issues_project_status_stats 
ON project_id, status_id FROM issues;

-- Regular maintenance
VACUUM ANALYZE issues;
REINDEX TABLE CONCURRENTLY issues;
```

---

## Schema Maintenance

### Backup Strategy

```sql
-- Daily incremental backups
pg_dump --format=custom --file=mytodo_$(date +%Y%m%d).backup mytodo_dev

-- Point-in-time recovery
-- Enable WAL archiving in postgresql.conf
```

### Monitoring Queries

```sql
-- Table sizes
SELECT 
    schemaname,
    tablename,
    pg_size_pretty(pg_total_relation_size(schemaname||'.'||tablename)) AS size
FROM pg_tables
WHERE schemaname = 'public'
ORDER BY pg_total_relation_size(schemaname||'.'||tablename) DESC;

-- Index usage
SELECT 
    schemaname,
    tablename,
    indexname,
    idx_scan,
    idx_tup_read,
    idx_tup_fetch
FROM pg_stat_user_indexes
ORDER BY idx_scan ASC;

-- Slow queries
SELECT 
    query,
    calls,
    mean_exec_time,
    total_exec_time
FROM pg_stat_statements
ORDER BY mean_exec_time DESC
LIMIT 20;
```

---

## Migration Plan

### Phase 1: Core Tables
1. users, sessions, oauth_providers
2. organizations, organization_members
3. projects, project_members

### Phase 2: Issue System
1. issue_types, issue_statuses, issue_priorities
2. issues (with partitioning)
3. issue_links, issue_labels, issue_watchers

### Phase 3: Collaboration
1. comments (with partitioning)
2. attachments
3. worklogs

### Phase 4: Advanced Features
1. boards, board_columns, sprints
2. workflows, workflow_transitions
3. automation_rules
4. notifications, integrations, webhooks

---

## Conclusion

This schema is production-ready and supports:

✅ **Scalability**: Partitioning for billions of records  
✅ **Performance**: Comprehensive indexing and denormalization  
✅ **Flexibility**: JSONB for custom fields  
✅ **Audit**: Complete history tracking  
✅ **Multi-tenancy**: Organization-based isolation  
✅ **Real-time**: Efficient queries for dashboards  
✅ **Integration**: Webhooks and external systems  

For implementation, see:
- [Migration Guide](migration-guide.md)
- [Query Optimization](queries/optimization.md)
- [ER Diagram](er-diagram.md)
