CREATE TABLE sessions (
    account_id VARCHAR(128) NOT NULL,
    session_id VARCHAR(256) NOT NULL UNIQUE,
    session_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    session_expires_at TIMESTAMP NOT NULL DEFAULT NOW() + "30 day"::INTERVAL,
    PRIMARY KEY (account_id, session_id),
    FOREIGN KEY (account_id)
        REFERENCES accounts(account_id)
        ON DELETE CASCADE
);
