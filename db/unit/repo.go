package unit

import (
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
