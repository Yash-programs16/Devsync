CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    workspace_id UUID NOT NULL
        REFERENCES workspaces(id) ON DELETE CASCADE,

    name VARCHAR(50) NOT NULL,
    description TEXT,

    is_default BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(workspace_id, name)
);

CREATE INDEX idx_roles_workspace_id
    ON roles(workspace_id);