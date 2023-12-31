// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package gen

import (
	"context"
)

const userCreate = `-- name: UserCreate :one
INSERT INTO users (id) VALUES (?) RETURNING id, created_at, updated_at, deleted_at
`

func (q *Queries) UserCreate(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, userCreate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const userDeletePermanently = `-- name: UserDeletePermanently :exec
DELETE FROM users
WHERE id = ?
`

func (q *Queries) UserDeletePermanently(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, userDeletePermanently, id)
	return err
}

const userDeleteSoftly = `-- name: UserDeleteSoftly :exec
UPDATE users
set deleted_at = CURRENT_TIMESTAMP
WHERE id = ?
`

func (q *Queries) UserDeleteSoftly(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, userDeleteSoftly, id)
	return err
}

const userEnsureExists = `-- name: UserEnsureExists :exec
INSERT INTO users (id) VALUES (?) ON CONFLICT (id) DO NOTHING
`

func (q *Queries) UserEnsureExists(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, userEnsureExists, id)
	return err
}

const userGet = `-- name: UserGet :one
SELECT id, created_at, updated_at, deleted_at FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) UserGet(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, userGet, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const userList = `-- name: UserList :many
SELECT id, created_at, updated_at, deleted_at FROM users
ORDER BY id
`

func (q *Queries) UserList(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, userList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
