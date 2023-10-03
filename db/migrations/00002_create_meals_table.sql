-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE meals (
    id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT    NOT NULL,

    author_id INTEGER NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME          DEFAULT NULL,

    FOREIGN KEY (author_id) REFERENCES users(id)
);

-- +goose Down
DROP TABLE meals;
