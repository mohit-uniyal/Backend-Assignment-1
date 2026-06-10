BEGIN;

DROP TRIGGER IF EXISTS update_table_timestamp ON events;
DROP TRIGGER IF EXISTS update_table_timestamp ON users;
DROP TRIGGER IF EXISTS update_table_timestamp ON reservations;

DROP FUNCTION IF EXISTS update_modified_column();

ALTER TABLE events
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;

ALTER TABLE users
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;

ALTER TABLE reservations
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;

COMMIT;