-- Update init.sql to include auth schema initialization

\c mytodo_dev

-- Import auth schema
\i /docker-entrypoint-initdb.d/init_auth_schema.sql

-- Create initial admin user (optional - for testing)
-- Password: admin123 (bcrypt hash)
-- INSERT INTO users (id, email, password_hash, name, is_active, created_at, updated_at)
-- VALUES (
--     'a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11'::uuid,
--     'admin@mytodo.com',
--     '$2a$10$rYvN8F7HZJZvEX.FJ0V0jOHVNvD1q8X8IlU2bXYK7GJvM5KJgK3cy',
--     'Admin User',
--     true,
--     NOW(),
--     NOW()
-- );

-- Grant necessary permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO mytodo;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO mytodo;
