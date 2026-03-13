-- Bootstrap Migration: Auth Base Tables
-- This is the foundational migration that must run before all others.
-- It creates the pgcrypto extension, the shared update_updated_at_column()
-- trigger function, the users table, and the sessions table.
--
-- All subsequent migrations that reference users(id) depend on this file.

-- Required for gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Shared trigger function used across all schema tables
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Users table: core auth credential store
CREATE TABLE IF NOT EXISTS users (
    id            UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name          VARCHAR(255) NOT NULL,
    is_active     BOOLEAN      NOT NULL DEFAULT true,
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP    NOT NULL DEFAULT NOW(),
    last_login_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_email     ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);

CREATE OR REPLACE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Sessions table: refresh-token and device tracking for login sessions
CREATE TABLE IF NOT EXISTS sessions (
    id            UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id       UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    refresh_token TEXT        NOT NULL UNIQUE,
    user_agent    TEXT,
    ip_address    VARCHAR(45),
    expires_at    TIMESTAMP   NOT NULL,
    created_at    TIMESTAMP   NOT NULL DEFAULT NOW(),
    revoked_at    TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_sessions_user_id       ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_refresh_token ON sessions(refresh_token);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at    ON sessions(expires_at);

-- Comments for clarity
COMMENT ON TABLE users    IS 'Core user accounts and authentication credentials';
COMMENT ON TABLE sessions IS 'Active user sessions keyed on refresh tokens';
COMMENT ON COLUMN sessions.revoked_at IS 'NULL means the session is still valid; set on logout or forced expiry';
