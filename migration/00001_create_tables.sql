-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "users" (
    "id" serial not null unique,
    "login" varchar(255) not null unique,
    "password_hash" varchar(255) not null,
    "age" int not null,
    "user_role" varchar(10) not null,
    "token" varchar(255) DEFAULT NULL,
    "deleted_token" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "unique_user_login" UNIQUE("login")
);

CREATE TABLE "directors" (
    "id" serial not null unique,
    "name" varchar(255) not null,
    "date_of_birth" varchar(255) not null,
    CONSTRAINT "unique_director_name" UNIQUE("name")
);

CREATE TABLE "films" (
    "id" serial not null unique,
    "name" varchar(255) not null,
    "genre" varchar(255) not null,
    "director_id" int references "directors" (id) not null,
    "rate" decimal not null,
    "year" int not null,
    "minutes" int not null,
    CONSTRAINT "unique_film_name" UNIQUE("name")
);

CREATE TABLE "favourite" (
    "id" serial not null unique,
    "user_id" int references "users" (id) not null,
    "film_id" int references "films" (id) not null,
    PRIMARY KEY("user_id", "film_id")
);

CREATE TABLE "wishlist" (
    "id" serial not null unique,
    "user_id" int references "users" (id) not null,
    "film_id" int references "films" (id) not null,
    PRIMARY KEY("user_id", "film_id")
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "wishlist";
DROP TABLE "favourite";
DROP TABLE "films";
DROP TABLE "directors";
DROP TABLE "users";