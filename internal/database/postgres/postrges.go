package postgres

import (
	"database/sql"
	"log"
)

func SetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://kamila:password@localhost/apidb?sslmode=disable")
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {

		db.Close()
		return nil, err
	}

	log.Println("ok ping")

	return db, nil
}
