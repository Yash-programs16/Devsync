CREATE TABLE workspace_invites (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    workspace_id UUID NOT NULL
        REFERENCES workspaces(id) ON DELETE CASCADE,

    inviter_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    email VARCHAR(255) NOT NULL,

    token_hash TEXT NOT NULL UNIQUE,

    expires_at TIMESTAMPTZ NOT NULL,
    accepted_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_workspace_invites_workspace_id
    ON workspace_invites(workspace_id);

CREATE INDEX idx_workspace_invites_email
    ON workspace_invites(email);

CREATE INDEX idx_workspace_invites_token_hash
    ON workspace_invites(token_hash);