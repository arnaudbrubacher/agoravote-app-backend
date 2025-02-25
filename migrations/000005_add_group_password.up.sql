ALTER TABLE public.groups
    ADD COLUMN password text,
    ADD COLUMN requires_password boolean DEFAULT false NOT NULL;