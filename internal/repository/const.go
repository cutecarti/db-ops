package repository

const (
	CREATE_QUERRY = "INSERT INTO records (name) VALUES ($1) RETURNING id"
	GET_QUERRY    = "SELECT id, name FROM records WHERE id = $1"
	DELETE_QUERRY = "DELETE FROM records WHERE id = $1"
	UPDATE_QUERRY = "UPDATE records SET name = $1 WHERE id = $2"
)
