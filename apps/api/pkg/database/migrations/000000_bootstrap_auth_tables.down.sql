-- Rollback: Bootstrap Auth Tables
-- WARNING: dropping users cascades to sessions and all tables with users(id) FK.
-- Only run this in a full database teardown (not a partial rollback).

DROP TRIGGER IF EXISTS update_users_updated_at ON users;

DROP INDEX IF EXISTS idx_sessions_expires_at;
DROP INDEX IF EXISTS idx_sessions_refresh_token;
DROP INDEX IF EXISTS idx_sessions_user_id;

DROP INDEX IF EXISTS idx_users_is_active;
DROP INDEX IF EXISTS idx_users_email;

DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS users;

DROP FUNCTION IF EXISTS update_updated_at_column();
