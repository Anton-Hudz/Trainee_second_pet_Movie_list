-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO director (name, date_of_birth)
VALUES ('James Cameron', '1954');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from director;