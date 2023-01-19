-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO director (name, date_of_birth)
VALUES ('James Cameron', '1954'),
('Steven Spielberg', '1946'),
('Quentin Tarantino', '1963'),
('Robert Zemeckis', '1952'),
('Frank Darabont', '1959'),
('Roman Polanski', '1933'),
('Ridley Scott', '1937'),
('Andrew Stanton', '1965');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from director;