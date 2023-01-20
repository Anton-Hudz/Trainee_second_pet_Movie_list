-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO films (name, genre, director_id, rate, year, minutes)
VALUES ('Avatar', 'fantasy', 1, 8.62, 2009, 162), 
('Titanic', 'drama', 1, 7.16, 2000, 120),
('Shindlers List', 'drama', 2, 8.9, 1993, 180),
('Forest Gump', 'drama', 4, 8.8, 1994, 142),
('Green Mile', 'drama', 5, 8.5, 1999, 189),
('Saving Private Ryan', 'action', 2, 8.6, 1999, 169),
('The Pianist', 'drama', 6, 8.5, 2002, 150),
('Gladiator', 'historical', 7, 8.5, 2000, 155),
('WALL-E', 'fantasy', 8, 8.4, 2008, 98);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE from wishlist;
DELETE from favourite;
DELETE from films;