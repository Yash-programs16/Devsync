CREATE TABLE email_verifications (

    id UUID PRIMARY KEY,

    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    email VARCHAR(255) NOT NULL,

    otp_hash VARCHAR(255) NOT NULL,

    expires_at TIMESTAMP NOT NULL,

    verified BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT NOW()

);


CREATE INDEX idx_email_verifications_user_id
ON email_verifications(user_id);


CREATE INDEX idx_email_verifications_email
ON email_verifications(email);