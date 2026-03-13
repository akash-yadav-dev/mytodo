-- Rollback: Issues schema

DROP TRIGGER IF EXISTS update_issues_updated_at ON issues;

DROP INDEX IF EXISTS idx_issues_assignee_id;
DROP INDEX IF EXISTS idx_issues_priority;
DROP INDEX IF EXISTS idx_issues_status;
DROP INDEX IF EXISTS idx_issues_project_id;

DROP TABLE IF EXISTS issues;
