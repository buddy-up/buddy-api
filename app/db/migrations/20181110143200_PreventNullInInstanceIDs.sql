
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE instanceids ALTER COLUMN android SET NOT NULL;
ALTER TABLE instanceids ALTER COLUMN ios SET NOT NULL;
ALTER TABLE instanceids ALTER COLUMN web SET NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE instanceids ALTER COLUMN web DROP  NOT NULL;
ALTER TABLE instanceids ALTER COLUMN web DROP NOT NULL;
ALTER TABLE instanceids ALTER COLUMN web DROP NOT NULL;

