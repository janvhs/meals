package meal

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

type CreateParams = gen.MealCreateParams

func (r *Repository) Create(ctx context.Context, params CreateParams) (Meal, error) {
	return r.queries.MealCreate(ctx, params)
}

func (r *Repository) List(ctx context.Context) ([]Meal, error) {
	return r.queries.MealList(ctx)
}
