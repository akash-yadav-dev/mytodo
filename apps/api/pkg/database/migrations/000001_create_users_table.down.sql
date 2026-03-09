-- Rollback users module migration

DROP TRIGGER IF EXISTS update_user_preferences_updated_at ON user_preferences;
DROP TRIGGER IF EXISTS update_user_profiles_updated_at ON user_profiles;

DROP INDEX IF EXISTS idx_user_preferences_user_id;
DROP INDEX IF EXISTS idx_user_profiles_username;
DROP INDEX IF EXISTS idx_user_profiles_user_id;

DROP TABLE IF EXISTS user_preferences;
DROP TABLE IF EXISTS user_profiles;
