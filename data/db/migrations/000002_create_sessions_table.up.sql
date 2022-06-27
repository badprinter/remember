CREATE TABLE IF NOT EXISTS sessions (
    user_id SERIAL REFERENCES users(id),
    login_time TIMESTAMP NOT NULL,
    alive TIME NOT NULL
    jwt VARCHAR(30) NOT NULL
);

