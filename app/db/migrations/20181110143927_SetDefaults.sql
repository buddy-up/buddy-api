
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE instanceids ALTER COLUMN android SET DEFAULT '';
ALTER TABLE instanceids ALTER COLUMN ios SET DEFAULT  '';
ALTER TABLE instanceids ALTER COLUMN web SET DEFAULT '';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE instanceids ALTER COLUMN web DROP  DEFAULT ;
ALTER TABLE instanceids ALTER COLUMN web DROP DEFAULT ;
ALTER TABLE instanceids ALTER COLUMN web DROP DEFAULT ;

