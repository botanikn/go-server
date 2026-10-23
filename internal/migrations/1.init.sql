CREATE SCHEMA public;

-- 1. Таблица пользователей
CREATE TABLE users (
    id           SERIAL PRIMARY KEY,
    username     VARCHAR(50)  NOT NULL UNIQUE,
    passwordHash VARCHAR(255) NOT NULL,
    email        VARCHAR(255) NOT NULL UNIQUE,
    role         VARCHAR(20)  NOT NULL DEFAULT 'user',

    CONSTRAINT users_role_check
        CHECK (role IN ('user', 'admin'))
);

-- 2. Таблица токенов
CREATE TABLE tokens (
    id      SERIAL PRIMARY KEY,
    token   VARCHAR(512) NOT NULL UNIQUE,
    exp     INTEGER      NOT NULL,
    user_id INTEGER      NOT NULL,

    CONSTRAINT fk_tokens_user
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
);

-- 3. Таблица заметок
CREATE TABLE notes (
    id      SERIAL PRIMARY KEY,
    title   VARCHAR(255) NOT NULL,
    content TEXT         NOT NULL,
    user_id INTEGER      NOT NULL,

    CONSTRAINT fk_notes_user
        FOREIGN KEY (user_id)
        REFERENCES users (id)
        ON DELETE CASCADE
);