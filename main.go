package main

import (
	"clean-arsitektur/simple-restapi/db"
	"clean-arsitektur/simple-restapi/routes"

	_ "database/sql"

	_ "github.com/lib/pq"
)

func init() {
	db.Config()
}

func main() {
	routes.Routes()
}
