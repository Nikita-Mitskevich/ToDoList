CREATE SCHEMA todoapp;

CREATE TABLE IF NOT EXISTS todoapp.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK (CHAR_LENGTH(full_name) BETWEEN 3 AND 100),
    phone VARCHAR(15) CHECK(CHAR_LENGTH(phone) BETWEEN 10 AND 15 AND phone LIKE '+%')
);

CREATE TABLE IF NOT EXISTS todoapp.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK(CHAR_LENGTH(title) BETWEEN 1 AND 100),
    description VARCHAR(1000) CHECK (CHAR_LENGTH(description) BETWEEN 1 AND 1000),
    completed BOOL NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL,
    completed_at TIMESTAMPTZ,
    author_user_id INTEGER NOT NULL REFERENCES todoapp.users(id),

    CHECK (
        (completed = FALSE AND completed_at IS NULL) OR (completed = TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    )
);