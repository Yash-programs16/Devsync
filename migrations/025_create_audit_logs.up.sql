CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    workspace_id UUID
        REFERENCES workspaces(id) ON DELETE CASCADE,

    user_id UUID
        REFERENCES users(id) ON DELETE SET NULL,

    action VARCHAR(100) NOT NULL,

    entity_type VARCHAR(50),

    entity_id UUID,

    metadata JSONB,

    ip_address INET,

    user_agent TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_audit_logs_workspace_id
    ON audit_logs(workspace_id);

CREATE INDEX idx_audit_logs_user_id
    ON audit_logs(user_id);

CREATE INDEX idx_audit_logs_entity
    ON audit_logs(entity_type, entity_id);

CREATE INDEX idx_audit_logs_created_at
    ON audit_logs(created_at);