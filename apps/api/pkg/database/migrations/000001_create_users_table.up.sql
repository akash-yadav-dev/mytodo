-- Migration 000001: User Profiles & Preferences
-- Extends the core users table with profile and notification preference tables.
-- Depends on: 000000_bootstrap_auth_tables (users table + update_updated_at_column fn)

-- User profiles table (extends auth.users)
CREATE TABLE IF NOT EXISTS user_profiles (
    id           UUID         PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID         NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    username     VARCHAR(50)  UNIQUE,
    display_name VARCHAR(100),
    avatar_url   TEXT,
    bio          TEXT,
    location     VARCHAR(100),
    website      VARCHAR(255),
    phone        VARCHAR(20),
    timezone     VARCHAR(50)  DEFAULT 'UTC',
    language     VARCHAR(10)  DEFAULT 'en',
    theme        VARCHAR(20)  DEFAULT 'light',
    created_at   TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- User preferences table
CREATE TABLE IF NOT EXISTS user_preferences (
    id                      UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id                 UUID      NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    email_notifications     BOOLEAN   DEFAULT true,
    push_notifications      BOOLEAN   DEFAULT true,
    newsletter_subscription BOOLEAN   DEFAULT false,
    weekly_digest           BOOLEAN   DEFAULT true,
    mentions_notifications  BOOLEAN   DEFAULT true,
    created_at              TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_user_profiles_user_id  ON user_profiles(user_id);
CREATE INDEX IF NOT EXISTS idx_user_profiles_username ON user_profiles(username);
CREATE INDEX IF NOT EXISTS idx_user_preferences_user_id ON user_preferences(user_id);

-- Triggers
CREATE OR REPLACE TRIGGER update_user_profiles_updated_at
    BEFORE UPDATE ON user_profiles
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

CREATE OR REPLACE TRIGGER update_user_preferences_updated_at
    BEFORE UPDATE ON user_preferences
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Comments
COMMENT ON TABLE user_profiles  IS 'Extended user profile information';
COMMENT ON TABLE user_preferences IS 'User notification and application preferences';
COMMENT ON COLUMN user_profiles.username     IS 'Unique username for the user (optional)';
COMMENT ON COLUMN user_profiles.display_name IS 'Display name shown in the application';
