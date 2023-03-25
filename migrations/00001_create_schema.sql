-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE social_user
(
    id              UUID primary key,
    first_name      text                    NOT NULL,
    second_name     text                    NOT NULL,
    age             integer                 NOT NULL,
    sex             int                     NOT NULL,
    city            text                    NOT NULL,
    biography       text                    NOT NULL,
    hashed_password text                    NOT NULL,
    created_at      timestamp default now() not null,
    updated_at      timestamp default now() not null
);

CREATE TABLE user_session
(
    id         UUID primary key,
    user_id    UUID                    NOT NULL,
    token      text                    NOT NULL,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE social_user;
DROP TABLE user_session;
