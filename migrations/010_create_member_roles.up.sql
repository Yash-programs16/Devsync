CREATE TABLE member_roles (
    member_id UUID NOT NULL
        REFERENCES workspace_members(id) ON DELETE CASCADE,

    role_id UUID NOT NULL
        REFERENCES roles(id) ON DELETE CASCADE,

    PRIMARY KEY (member_id, role_id)
);

CREATE INDEX idx_member_roles_role_id
    ON member_roles(role_id);


    