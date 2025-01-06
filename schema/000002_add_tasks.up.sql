CREATE TABLE IF NOT EXISTS tasks
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) DEFAULT null,
    price int DEFAULT 1
);

CREATE TABLE IF NOT EXISTS task_complete
(
    id SERIAL PRIMARY KEY,
    user_id int references users (id) on delete cascade not null,
    task_id int references tasks (id) on delete cascade not null
);

