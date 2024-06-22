package database

import (
	"database/sql"
	"log"
	"simple-go-blog-webserver/internal/models"

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

func CreateUser(newUser models.User) error {
	_, err := db.Exec("insert into blog_webserver_db.users (name, username, email) values (?, ?, ?)", newUser.Name, newUser.UserName, newUser.Email)
	return err
}
