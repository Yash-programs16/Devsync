CREATE TABLE notification_settings (
    user_id UUID PRIMARY KEY
        REFERENCES users(id) ON DELETE CASCADE,

    email_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    mention_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    message_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    workspace_invite_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);