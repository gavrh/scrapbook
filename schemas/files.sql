CREATE TABLE files (
    account_id VARCHAR(128) NOT NULL,
    file_name VARCHAR(256) NOT NULL,
    file_type VARCHAR(24) NOT NULL,
    file_updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    file_created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, file_name)
);
