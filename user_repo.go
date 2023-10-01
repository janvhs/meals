package main

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) UserByID(id string) (User, error) {
	var user User

	err := r.DB.Get(&user, "SELECT * FROM users WHERE id = $1 LIMIT 1", id)

	return user, err
}

func (r *UserRepository) EnsureExists(id string) error {
	_, err := r.DB.Exec("INSERT INTO users (id) VALUES ($1) ON CONFLICT (id) DO NOTHING", id)

	return err
}
