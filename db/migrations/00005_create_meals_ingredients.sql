-- SQL-Dialect: SQLite

-- +goose Up
CREATE TABLE meals_ingredients (
    meal_id       INTEGER NOT NULL,
    ingredient_id INTEGER NOT NULL,
    amount        INTEGER NOT NULL,
    unit_id       INTEGER NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME          DEFAULT NULL,

    FOREIGN KEY (meal_id)       REFERENCES meals(id)
    FOREIGN KEY (ingredient_id) REFERENCES ingredients(id)
    FOREIGN KEY (unit_id)       REFERENCES units(id)
);

-- +goose Down
DROP TABLE meals_ingredients;
