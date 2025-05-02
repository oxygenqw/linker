CREATE TABLE students (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    user_name VARCHAR(50),
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    university VARCHAR(100),
    faculty VARCHAR(100),
    idea TEXT,
    about TEXT,
    github VARCHAR(100),
    job VARCHAR(100),
    course VARCHAR(50),
    education VARCHAR(50)
);

CREATE TABLE teachers (
    id uuid PRIMARY KEY,
    telegram_id BIGINT,
    user_name VARCHAR(50),
    first_name VARCHAR(50),
    middle_name VARCHAR(50),
    last_name VARCHAR(50),
    university VARCHAR(100),
    faculty VARCHAR(100),
    idea TEXT,
    about TEXT,
    degree VARCHAR(50),
    position VARCHAR(50),
    is_free BOOLEAN
);