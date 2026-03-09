-- Seed data for users module
-- This script populates sample data for development and testing

-- Note: Users are created via the auth module
-- These are seed profiles that reference auth users

-- Sample user profiles (assuming auth users exist)
INSERT INTO user_profiles (id, user_id, username, display_name, avatar_url, bio, location, timezone, language, theme)
VALUES 
    (gen_random_uuid(), 
     (SELECT id FROM users WHERE email = 'admin@example.com' LIMIT 1),
     'admin',
     'System Administrator',
     'https://api.dicebear.com/7.x/avataaars/svg?seed=admin',
     'System administrator and maintainer',
     'San Francisco, CA',
     'America/Los_Angeles',
     'en',
     'light'
    ),
    (gen_random_uuid(),
     (SELECT id FROM users WHERE email = 'john.doe@example.com' LIMIT 1),
     'johndoe',
     'John Doe',
     'https://api.dicebear.com/7.x/avataaars/svg?seed=john',
     'Software engineer passionate about building great products',
     'New York, NY',
     'America/New_York',
     'en',
     'dark'
    ),
    (gen_random_uuid(),
     (SELECT id FROM users WHERE email = 'jane.smith@example.com' LIMIT 1),
     'janesmith',
     'Jane Smith',
     'https://api.dicebear.com/7.x/avataaars/svg?seed=jane',
     'Product manager with a focus on user experience',
     'London, UK',
     'Europe/London',
     'en',
     'light'
    )
ON CONFLICT (user_id) DO NOTHING;

-- Sample user preferences
INSERT INTO user_preferences (id, user_id, email_notifications, push_notifications, newsletter_subscription, weekly_digest, mentions_notifications)
SELECT 
    gen_random_uuid(),
    user_id,
    true,
    true,
    false,
    true,
    true
FROM user_profiles
ON CONFLICT (user_id) DO NOTHING;

-- Display statistics
SELECT 
    (SELECT COUNT(*) FROM user_profiles) as total_profiles,
    (SELECT COUNT(*) FROM user_preferences) as total_preferences;
