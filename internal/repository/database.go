package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	connection *sql.DB
}

func (db *Database) OpenConnection() (err error) {
	db.connection, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/blog_webserver_db")
	return
}

func (db *Database) CloseConnection() (err error) {
	return db.connection.Close()
}

func (db *Database) CheckConnection() error {
	err := db.connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
