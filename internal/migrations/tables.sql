CREATE TABLE students (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50)
);

CREATE TABLE teachers (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50)
);
