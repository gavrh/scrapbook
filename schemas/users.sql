CREATE TABLE users (
    account_id VARCHAR(128) NOT NULL UNIQUE,
    user_login VARCHAR(32) NOT NULL UNIQUE,
    user_password VARCHAR(128) NOT NULL,
    user_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, user_login),
    FOREIGN KEY (account_id)
        REFERENCES accounts(account_id)
        ON DELETE CASCADE
);
