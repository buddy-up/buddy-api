
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE users
 ADD PRIMARY KEY(id) ;

CREATE TABLE InstanceIds (
  instanceId varchar(255),
  id numeric,
  PRIMARY KEY (InstanceId),
  FOREIGN KEY (id) REFERENCES users(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE users DROP CONSTRAINT id_pkey;
DROP TABLE InstanceIds;
