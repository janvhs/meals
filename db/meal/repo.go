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

type MealCreateParams = gen.MealCreateParams

func (r *Repository) Create(ctx context.Context, params MealCreateParams) (Meal, error) {
	return r.queries.MealCreate(ctx, params)
}

func (r *Repository) List(ctx context.Context, id string) ([]Meal, error) {
	return r.queries.MealList(ctx)
}
