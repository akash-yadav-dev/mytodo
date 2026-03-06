Below is a **production-grade relational schema (≈60 tables)** for a **Jira-style task management / todo platform**, designed for a **Go (Golang) backend + Next.js frontend**.

The design assumes:

* **Database:** PostgreSQL
* **Backend:** Go (Clean Architecture / Hexagonal)
* **Frontend:** Next.js (App Router)
* **ORM:** `sqlc` or `ent` (recommended for Go)
* **Auth:** JWT + OAuth
* **Search:** ElasticSearch / Meilisearch
* **Queue:** Redis / NATS
* **File storage:** S3 compatible storage
* **Cache:** Redis

This schema supports:

* organizations & multi-tenancy
* projects
* issues / tasks
* subtasks
* sprints / boards
* comments
* attachments
* notifications
* activity logs
* permissions
* workflows
* automation rules
* integrations

---

# 1. Core Infrastructure

| Table                 | Purpose             |
| --------------------- | ------------------- |
| organizations         | Workspace / company |
| organization_settings | Org configuration   |
| organization_domains  | Domain verification |
| organization_invites  | Invite users        |

### organizations

```
id (uuid) PK
name
slug
plan_id
created_by
created_at
```

### organization_settings

```
id
organization_id
default_timezone
default_language
created_at
```

### organization_domains

```
id
organization_id
domain
verified
created_at
```

### organization_invites

```
id
organization_id
email
role_id
token
expires_at
```

---

# 2. User & Authentication

| Table          | Purpose           |
| -------------- | ----------------- |
| users          | Global users      |
| user_profiles  | Profile metadata  |
| user_settings  | User preferences  |
| user_sessions  | Login sessions    |
| oauth_accounts | OAuth providers   |
| user_emails    | Multiple emails   |
| user_activity  | Activity tracking |

### users

```
id
email
password_hash
is_active
created_at
```

### user_profiles

```
id
user_id
name
avatar_url
bio
```

### user_settings

```
user_id
theme
timezone
language
```

### user_sessions

```
id
user_id
token
ip_address
expires_at
```

### oauth_accounts

```
id
user_id
provider
provider_user_id
access_token
```

### user_emails

```
id
user_id
email
verified
primary_email
```

### user_activity

```
id
user_id
last_seen
last_ip
```

---

# 3. Roles & Permissions

| Table                | Purpose                |
| -------------------- | ---------------------- |
| roles                | Role definitions       |
| permissions          | Permission list        |
| role_permissions     | Mapping                |
| organization_members | Users inside workspace |
| project_members      | Users inside project   |

### roles

```
id
name
scope (organization/project)
```

### permissions

```
id
name
description
```

### role_permissions

```
role_id
permission_id
```

### organization_members

```
id
organization_id
user_id
role_id
joined_at
```

### project_members

```
id
project_id
user_id
role_id
```

---

# 4. Projects

| Table              | Purpose               |
| ------------------ | --------------------- |
| projects           | Project container     |
| project_settings   | Project configuration |
| project_labels     | Labels                |
| project_versions   | Releases              |
| project_components | Components/modules    |
| project_categories | Project grouping      |

### projects

```
id
organization_id
name
key
description
created_by
created_at
```

### project_settings

```
project_id
issue_prefix
default_assignee
```

### project_labels

```
id
project_id
name
color
```

### project_versions

```
id
project_id
name
release_date
status
```

### project_components

```
id
project_id
name
lead_user_id
```

### project_categories

```
id
organization_id
name
```

---

# 5. Boards & Sprints

| Table               | Purpose             |
| ------------------- | ------------------- |
| boards              | Scrum/Kanban boards |
| board_columns       | Columns             |
| board_column_status | Mapping statuses    |
| sprints             | Sprint info         |
| sprint_issues       | Issues in sprint    |

### boards

```
id
project_id
name
type (scrum/kanban)
created_at
```

### board_columns

```
id
board_id
name
position
```

### board_column_status

```
column_id
status_id
```

### sprints

```
id
board_id
name
goal
start_date
end_date
state
```

### sprint_issues

```
sprint_id
issue_id
```

---

# 6. Issues / Tasks (Core)

| Table            | Purpose          |
| ---------------- | ---------------- |
| issues           | Main task entity |
| issue_types      | Bug, task, story |
| issue_status     | Workflow states  |
| issue_priorities | Priority         |
| issue_links      | Dependencies     |
| issue_labels     | Label mapping    |
| issue_watchers   | Watch list       |

### issues

```
id
project_id
title
description
issue_type_id
status_id
priority_id
assignee_id
reporter_id
parent_issue_id
story_points
due_date
created_at
updated_at
```

### issue_types

```
id
name
icon
```

### issue_status

```
id
name
category
```

### issue_priorities

```
id
name
level
color
```

### issue_links

```
id
source_issue_id
target_issue_id
type
```

### issue_labels

```
issue_id
label_id
```

### issue_watchers

```
issue_id
user_id
```

---

# 7. Comments & Activity

| Table             | Purpose          |
| ----------------- | ---------------- |
| comments          | Issue comments   |
| comment_reactions | Emoji reactions  |
| activity_logs     | Activity history |
| issue_history     | Field changes    |

### comments

```
id
issue_id
user_id
content
created_at
updated_at
```

### comment_reactions

```
id
comment_id
user_id
emoji
```

### activity_logs

```
id
organization_id
actor_user_id
action
entity_type
entity_id
created_at
```

### issue_history

```
id
issue_id
field_name
old_value
new_value
changed_by
changed_at
```

---

# 8. Attachments & Files

| Table               | Purpose          |
| ------------------- | ---------------- |
| attachments         | Files            |
| attachment_versions | File versioning  |
| file_storage        | Storage metadata |

### attachments

```
id
issue_id
uploaded_by
file_name
file_size
storage_id
created_at
```

### attachment_versions

```
id
attachment_id
version
storage_id
created_at
```

### file_storage

```
id
provider
bucket
path
```

---

# 9. Notifications

| Table                 | Purpose             |
| --------------------- | ------------------- |
| notifications         | Notification events |
| user_notifications    | Delivery            |
| notification_settings | Preferences         |

### notifications

```
id
type
entity_type
entity_id
created_at
```

### user_notifications

```
id
notification_id
user_id
is_read
delivered_at
```

### notification_settings

```
user_id
email_enabled
push_enabled
```

---

# 10. Workflow Engine

| Table                | Purpose             |
| -------------------- | ------------------- |
| workflows            | Workflow            |
| workflow_states      | States              |
| workflow_transitions | Allowed transitions |
| workflow_rules       | Validation          |

### workflows

```
id
project_id
name
```

### workflow_states

```
id
workflow_id
name
category
```

### workflow_transitions

```
id
workflow_id
from_state
to_state
```

### workflow_rules

```
id
transition_id
rule_type
rule_config
```

---

# 11. Automation

| Table               | Purpose        |
| ------------------- | -------------- |
| automation_rules    | Automation     |
| automation_triggers | Trigger events |
| automation_actions  | Actions        |
| automation_logs     | Execution logs |

### automation_rules

```
id
project_id
name
enabled
```

### automation_triggers

```
id
rule_id
event_type
conditions
```

### automation_actions

```
id
rule_id
action_type
config
```

### automation_logs

```
id
rule_id
status
executed_at
```

---

# 12. Integrations

| Table              | Purpose               |
| ------------------ | --------------------- |
| integrations       | External integrations |
| integration_tokens | API tokens            |
| webhooks           | Webhooks              |
| webhook_deliveries | Delivery logs         |

### integrations

```
id
organization_id
provider
config
```

### integration_tokens

```
id
integration_id
token
expires_at
```

### webhooks

```
id
organization_id
url
secret
```

### webhook_deliveries

```
id
webhook_id
status
response_code
created_at
```

---

# Final Table Count

| Module        | Tables |
| ------------- | ------ |
| Organizations | 4      |
| Users/Auth    | 7      |
| Roles         | 5      |
| Projects      | 6      |
| Boards        | 5      |
| Issues        | 7      |
| Comments      | 4      |
| Files         | 3      |
| Notifications | 3      |
| Workflow      | 4      |
| Automation    | 4      |
| Integrations  | 4      |

**Total: ~56–60 tables**

---

# Recommended Tech Stack

## Backend (Go)

Recommended stack:

```
Go
Gin / Fiber
PostgreSQL
Redis
sqlc or ent ORM
NATS / Kafka
ElasticSearch
```

Folder structure:

```
backend
 ├── cmd
 │   └── api
 ├── internal
 │   ├── auth
 │   ├── organization
 │   ├── project
 │   ├── issue
 │   ├── sprint
 │   ├── comment
 │   ├── notification
 │   └── automation
 ├── pkg
 │   ├── database
 │   ├── logger
 │   └── middleware
```

---

## Frontend (Next.js)

```
Next.js 15
React Server Components
Tanstack Query
Zustand
Tailwind
Shadcn UI
```

Folder architecture:

```
frontend
 ├── app
 │   ├── dashboard
 │   ├── projects
 │   ├── boards
 │   └── issues
 ├── components
 ├── hooks
 ├── services
 ├── store
 └── lib
```

---

# DevOps Stack

| Tool           | Use              |
| -------------- | ---------------- |
| Docker         | containerization |
| Kubernetes     | orchestration    |
| Terraform      | infra            |
| GitHub Actions | CI/CD            |
| Prometheus     | metrics          |
| Grafana        | monitoring       |
| Sentry         | error tracking   |

---

# Scaling Design (like Jira)

Key architecture patterns:

```
API Gateway
     │
Go Microservices
 ├── Auth Service
 ├── Issue Service
 ├── Notification Service
 ├── Search Service
 └── Automation Service
     │
PostgreSQL (Primary DB)
Redis (Cache)
ElasticSearch (Search)
S3 (Files)
```

---

If you want, I can also show you:

* **A production-grade folder architecture (100+ folders) used by companies**
* **Exact Jira-style issue table design (very advanced)**
* **Complete system architecture diagram**
* **How to build this as a SaaS product (multi-tenant)**

Those will make this **a real startup-level system design.**
