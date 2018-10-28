
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users
  ADD COLUMN origin varchar(20);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE users
  DROP COLUMN origin;