CREATE TABLE message_reactions (
    message_id UUID NOT NULL
        REFERENCES messages(id) ON DELETE CASCADE,

    user_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    emoji VARCHAR(50) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (message_id, user_id, emoji)
);

CREATE INDEX idx_message_reactions_user_id
    ON message_reactions(user_id);