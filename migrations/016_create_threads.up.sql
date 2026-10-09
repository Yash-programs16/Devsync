CREATE TABLE threads (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    parent_message_id UUID NOT NULL
        REFERENCES messages(id) ON DELETE CASCADE,

    created_by UUID
        REFERENCES users(id) ON DELETE SET NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_threads_parent_message_id
    ON threads(parent_message_id);

CREATE INDEX idx_threads_created_by
    ON threads(created_by);