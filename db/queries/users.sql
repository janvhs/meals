-- name: UserGet :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: UserList :many
SELECT * FROM users
ORDER BY id;

-- name: UserCreate :one
INSERT INTO users (id) VALUES (?) RETURNING *;

-- name: UserEnsureExists :exec
INSERT INTO users (id) VALUES (?) ON CONFLICT (id) DO NOTHING;

-- name: UserDeletePermanently :exec
DELETE FROM users
WHERE id = ?;

-- name: UserDeleteSoftly :exec
UPDATE users
set deleted_at = CURRENT_TIMESTAMP
WHERE id = ?;
