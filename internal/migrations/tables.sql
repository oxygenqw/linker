CREATE TABLE users (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    sure_name VARCHAR(50)
);
