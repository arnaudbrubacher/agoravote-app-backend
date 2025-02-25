CREATE TABLE group_invitations (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    group_id uuid NOT NULL REFERENCES groups(id),
    email text NOT NULL,
    token text NOT NULL UNIQUE,
    expires_at timestamp with time zone NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    used boolean DEFAULT false NOT NULL
);

CREATE INDEX idx_group_invitations_token ON group_invitations(token);
CREATE INDEX idx_group_invitations_email ON group_invitations(email);