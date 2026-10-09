CREATE TABLE channels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    workspace_id UUID NOT NULL
        REFERENCES workspaces(id) ON DELETE CASCADE,

    name VARCHAR(100) NOT NULL,

    description TEXT,

    type VARCHAR(20) NOT NULL DEFAULT 'text',

    is_private BOOLEAN NOT NULL DEFAULT FALSE,

    created_by UUID REFERENCES users(id) ON DELETE SET NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(workspace_id, name),

    CHECK (type IN ('text', 'announcement', 'code'))
);

CREATE INDEX idx_channels_workspace_id
    ON channels(workspace_id);