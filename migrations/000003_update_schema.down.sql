-- Revert votes changes
ALTER TABLE public.votes
    ALTER COLUMN value DROP NOT NULL;

-- Revert posts changes
ALTER TABLE public.posts
    DROP COLUMN IF EXISTS title;

-- Revert group_members changes
ALTER TABLE public.group_members
    DROP CONSTRAINT IF EXISTS group_members_unique_membership,
    ALTER COLUMN is_admin DROP NOT NULL;

-- Revert users changes
ALTER TABLE public.users
    ALTER COLUMN email DROP NOT NULL,
    ALTER COLUMN name DROP NOT NULL;