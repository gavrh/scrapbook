CREATE TABLE accounts (
    account_id VARCHAR(128) NOT NULL UNIQUE,
    account_2fa_secret VARCHAR(16) NOT NULL UNIQUE,
    account_2fa_is_setup BOOL NOT NULL DEFAULT FALSE,
);
