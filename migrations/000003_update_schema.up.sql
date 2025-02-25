-- Make email and name NOT NULL in users table
ALTER TABLE public.users
    ALTER COLUMN email SET NOT NULL,
    ALTER COLUMN name SET NOT NULL;

-- Update group_members table
ALTER TABLE public.group_members
    ALTER COLUMN is_admin SET NOT NULL,
    ALTER COLUMN is_admin SET DEFAULT false;

-- Add unique constraint to prevent duplicate memberships
ALTER TABLE public.group_members
    ADD CONSTRAINT group_members_unique_membership UNIQUE (group_id, user_id);

-- Add title column to posts
ALTER TABLE public.posts
    ADD COLUMN title varchar(255) NOT NULL DEFAULT '';

-- Update votes table
ALTER TABLE public.votes
    ALTER COLUMN value SET NOT NULL;