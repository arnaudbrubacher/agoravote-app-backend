-- Update groups table
ALTER TABLE public.groups
    ALTER COLUMN name TYPE VARCHAR(255),
    ALTER COLUMN last_active TYPE TIMESTAMP WITH TIME ZONE 
        USING last_active::timestamp with time zone,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW();

-- Update group_members table
ALTER TABLE public.group_members
    RENAME COLUMN role TO is_admin;