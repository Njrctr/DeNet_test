CREATE TABLE IF NOT EXIST users
(
    id SERIAL PRIMARY KEY not null,
    name VARCHAR(255) not null unique,
    password_hash VARCHAR(255) not null
);