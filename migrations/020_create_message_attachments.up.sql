CREATE TABLE message_attachments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    message_id UUID NOT NULL
        REFERENCES messages(id) ON DELETE CASCADE,

    file_id UUID NOT NULL
        REFERENCES files(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(message_id, file_id)
);

CREATE INDEX idx_message_attachments_message_id
    ON message_attachments(message_id);

CREATE INDEX idx_message_attachments_file_id
    ON message_attachments(file_id);