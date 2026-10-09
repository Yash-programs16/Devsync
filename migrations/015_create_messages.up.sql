CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    channel_id UUID NOT NULL
        REFERENCES channels(id) ON DELETE CASCADE,

    sender_id UUID
        REFERENCES users(id) ON DELETE SET NULL,

    content TEXT NOT NULL,

    message_type VARCHAR(20) NOT NULL DEFAULT 'text',

    is_edited BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CHECK (
        message_type IN ('text', 'code', 'system')
    )
);

CREATE INDEX idx_messages_channel_id
    ON messages(channel_id);

CREATE INDEX idx_messages_sender_id
    ON messages(sender_id);

CREATE INDEX idx_messages_created_at
    ON messages(created_at);

ALTER TABLE channel_pins
ADD CONSTRAINT fk_channel_pins_message
FOREIGN KEY (message_id)
REFERENCES messages(id)
ON DELETE CASCADE;