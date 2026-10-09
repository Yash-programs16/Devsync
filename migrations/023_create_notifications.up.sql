CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        REFERENCES users(id) ON DELETE CASCADE,

    type VARCHAR(50) NOT NULL,

    title VARCHAR(255) NOT NULL,

    body TEXT,

    reference_id UUID,

    is_read BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_notifications_user_id
    ON notifications(user_id);

CREATE INDEX idx_notifications_user_unread
    ON notifications(user_id, is_read);

CREATE INDEX idx_notifications_created_at
    ON notifications(created_at);