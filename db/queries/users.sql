-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (id) VALUES (?) RETURNING *;

-- name: EnsureExistsUser :exec
INSERT INTO users (id) VALUES (?) ON CONFLICT (id) DO NOTHING;

-- name: DeleteUserPermanently :exec
DELETE FROM users
WHERE id = ?;

-- name: DeleteUserSoftly :exec
UPDATE users
set deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;