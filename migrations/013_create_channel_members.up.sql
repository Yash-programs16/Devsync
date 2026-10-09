CREATE TABLE channel_members (
    channel_id UUID NOT NULL
        REFERENCES channels(id) ON DELETE CASCADE,

    user_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (channel_id, user_id)
);

CREATE INDEX idx_channel_members_user_id
    ON channel_members(user_id);