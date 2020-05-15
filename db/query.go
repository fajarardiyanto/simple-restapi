package db

const (
	InsertTodoQuery = `INSERT INTO todos(id, title, description) VALUES(DEFAULT, $1, $2)`
	GetAllTodoQuery = `SELECT * FROM todos`
	GetTodoQuery    = `SELECT * FROM todos WHERE id = $1`
	UpdateTodoQuery = `UPDATE todos SET title = $1, description = $2 WHERE id = $3`
	DeleteQuery     = `DELETE FROM todos WHERE id = $1`
)
