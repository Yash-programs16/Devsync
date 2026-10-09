CREATE TABLE channel_pins (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    channel_id UUID NOT NULL
        REFERENCES channels(id) ON DELETE CASCADE,

    message_id UUID NOT NULL,

    pinned_by UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    pinned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(channel_id, message_id)
);

CREATE INDEX idx_channel_pins_channel_id
    ON channel_pins(channel_id);