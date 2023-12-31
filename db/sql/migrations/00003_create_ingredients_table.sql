-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE ingredients (
    id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL UNIQUE,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME          DEFAULT NULL
);

-- +goose Down
DROP TABLE ingredients;
