CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY not null,
    username VARCHAR(255) not null unique,
    balance INT not null DEFAULT 0,
    password_hash VARCHAR(255) not null,
    refer_from VARCHAR(255) not null DEFAULT "",
    refer_code VARCHAR(255) not null
);