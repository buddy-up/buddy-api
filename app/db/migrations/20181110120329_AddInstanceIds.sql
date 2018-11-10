
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE instanceids
  ADD COLUMN android varchar(255),
  ADD COLUMN ios varchar(255),
  ADD COLUMN web varchar(255),
  ADD  PRIMARY KEY(id),
  DROP COLUMN instanceid;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE instanceids
  DROP COLUMN android,
  DROP COLUMN ios,
  DROP COLUMN web,
  DROP CONSTRAINT id_pkey,
  ADD COLUMN instanceid varchar(255);