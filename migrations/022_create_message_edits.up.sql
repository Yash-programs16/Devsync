CREATE TABLE message_edits (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    message_id UUID NOT NULL
        REFERENCES messages(id) ON DELETE CASCADE,

    edited_by UUID
        REFERENCES users(id) ON DELETE SET NULL,

    previous_content TEXT NOT NULL,

    edited_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_message_edits_message_id
    ON message_edits(message_id);

CREATE INDEX idx_message_edits_edited_at
    ON message_edits(edited_at);