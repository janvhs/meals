-- name: MealList :many
SELECT * FROM meals
ORDER BY id;

-- name: MealCreate :one
INSERT INTO meals (
    name,
    author_id
) VALUES (
    ?,
    ?
) RETURNING *;
