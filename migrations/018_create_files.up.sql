CREATE TABLE files (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    workspace_id UUID NOT NULL
        REFERENCES workspaces(id) ON DELETE CASCADE,

    uploader_id UUID
        REFERENCES users(id) ON DELETE SET NULL,

    original_name VARCHAR(255) NOT NULL,

    storage_key TEXT NOT NULL UNIQUE,

    mime_type VARCHAR(255),

    size_bytes BIGINT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (size_bytes >= 0)
);

CREATE INDEX idx_files_workspace_id
    ON files(workspace_id);

CREATE INDEX idx_files_uploader_id
    ON files(uploader_id);