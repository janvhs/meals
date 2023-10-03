package user

import (
	"context"
	"database/sql"

	"git.bode.fun/meals/db/gen"
)

type Repository struct {
	queries *gen.Queries
}

func New(db *sql.DB) *Repository {
	return &Repository{
		queries: gen.New(db),
	}
}

func (r *Repository) Get(ctx context.Context, id string) (User, error) {
	return r.queries.UserGet(ctx, id)
}

func (r *Repository) EnsureExists(ctx context.Context, id string) error {
	return r.queries.UserEnsureExists(ctx, id)
}
