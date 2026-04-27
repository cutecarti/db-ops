package models

type Record struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
