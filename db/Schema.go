package db

func TodoSchema() {
	DB.Query(`
		CREATE TABLE todos (
			id SERIAL PRIMARY KEY,
			title TEXT,
			description TEXT
		)
	`)
}
