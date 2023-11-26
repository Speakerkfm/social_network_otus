-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE post
(
    id             UUID primary key DEFAULT uuid_generate_v1(),
    author_user_id UUID             DEFAULT uuid_generate_v1() NOT NULL,
    text           text                                        NOT NULL,
    created_at     timestamp        default now()              not null,
    updated_at     timestamp        default now()              not null
);


CREATE TABLE user_friend
(
    id        SERIAL PRIMARY KEY,
    user_id   UUID DEFAULT uuid_generate_v1() NOT NULL,
    friend_id UUID DEFAULT uuid_generate_v1() NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE post;
DROP TABLE user_friend;