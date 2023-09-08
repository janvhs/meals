-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE users (
    id TEXT NOT NULL PRIMARY KEY,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME          DEFAULT NULL
);

-- +goose Down
DROP TABLE users;
