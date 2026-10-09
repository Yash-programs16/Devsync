CREATE TABLE message_mentions (
    message_id UUID NOT NULL
        REFERENCES messages(id) ON DELETE CASCADE,

    user_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (message_id, user_id)
);

CREATE INDEX idx_message_mentions_user_id
    ON message_mentions(user_id);