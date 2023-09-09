package main

type User struct {
	BaseModel

	ID string `db:"id"`
}
