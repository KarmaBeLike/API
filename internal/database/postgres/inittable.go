package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY + login + password)")
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'users' was created successfully")
	return nil
}

func InsertUsers(login, password string) error {
	return nil
}
