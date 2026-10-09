CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT
);

CREATE INDEX idx_permissions_name
    ON permissions(name);