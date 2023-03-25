-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE EXTENSION pg_trgm;
CREATE INDEX social_user_first_name_idx ON social_user USING gin (first_name gin_trgm_ops);
CREATE INDEX social_user_second_name_idx ON social_user USING gin (first_name gin_trgm_ops);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX social_user_first_name_idx;
DROP INDEX social_user_second_name_idx;
