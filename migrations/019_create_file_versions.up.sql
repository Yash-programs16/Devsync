CREATE TABLE file_versions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    file_id UUID NOT NULL
        REFERENCES files(id) ON DELETE CASCADE,

    version_number INTEGER NOT NULL,

    storage_key TEXT NOT NULL,

    size_bytes BIGINT NOT NULL,

    uploaded_by UUID
        REFERENCES users(id) ON DELETE SET NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(file_id, version_number),

    CHECK (version_number > 0),
    CHECK (size_bytes >= 0)
);

CREATE INDEX idx_file_versions_file_id
    ON file_versions(file_id);