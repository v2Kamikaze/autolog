package database

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open(
		"pgx",
		"host=localhost port=5432 user=autolog password=autolog dbname=autolog sslmode=disable",
	)

	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	return db
}
