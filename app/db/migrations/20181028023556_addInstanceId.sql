
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE InstanceIds (
  instanceId varchar(255),
  id numeric,
  PRIMARY KEY (InstanceId),
  FOREIGN KEY (id) REFERENCES users(id)
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE InstanceIds;
