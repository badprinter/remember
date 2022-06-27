CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username VARCHAR(30) UNIQUE NOT NULL,
    password_encrypted VARCHAR(300) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
    id BIGSERIAL PRIMARY KEY,
    user_id serial REFERENCES users(id),
    document text
);