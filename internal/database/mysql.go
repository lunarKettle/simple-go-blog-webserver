package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func OpenConnection() (err error) {
	db, err = sql.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/blog_webserver_db")
	return
}

func CloseConnection() (err error) {
	return db.Close()
}

func CheckConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func CreateUser(name string, userName string, email string) error {
	_, err := db.Exec("insert into blog_webserver_db.users (name, username, email) values (?, ?, ?)", name, userName, email)
	return err
}
