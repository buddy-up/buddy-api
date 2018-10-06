
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users
  ALTER COLUMN id TYPE numeric;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
  ALTER COLUMN id TYPE int;