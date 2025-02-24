ALTER TABLE public.groups
    ALTER COLUMN name TYPE text,
    ALTER COLUMN last_active TYPE text,
    DROP COLUMN IF EXISTS created_at,
    DROP COLUMN IF EXISTS updated_at;

ALTER TABLE public.group_members
    RENAME COLUMN is_admin TO role;