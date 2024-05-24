-- +goose Up
CREATE TABLE accounts (
    id varchar(500) primary key,
    email varchar(255) unique not null,
    refresh_token varchar(500),
    enabled boolean not null default true
);

CREATE INDEX accounts_email_idx ON accounts(email);

-- +goose Down
DROP TABLE accounts;
