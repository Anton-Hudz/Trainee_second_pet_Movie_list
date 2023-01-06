-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "users" (
    "id" serial not null unique,
    "login" varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "age" int not null,
    "admin" varchar(10) DEFAULT NULL,
    "token" varchar(255) DEFAULT NULL,
    "deleted_token" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    CONSTRAINT "unique_user_login" UNIQUE("login")
);

CREATE TABLE "director" (
    "id" serial not null unique,
    "name" varchar(255) not null,
    "date_of_birth" varchar(255) not null
);

CREATE TABLE "film" (
    "id" serial not null unique,
    "name" varchar(255) not null,
    "genre" varchar(255) not null,
    "director_id" int references "director" (id) not null,
    "rate" int not null,
    "year" int not null,
    "minutes" int not null
);

CREATE TABLE "favourite" (
    "id" serial not null unique,
    "user_id" int references "users" (id),
    "film_id" int references "film" (id)
);

CREATE TABLE "wishlist" (
    "id" serial not null unique,
    "user_id" int references "users" (id),
    "film_id" int references "film" (id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "wishlist";
DROP TABLE "favourite";
DROP TABLE "film";
DROP TABLE "director";
DROP TABLE "users";