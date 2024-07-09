package database

import (
	"database/sql"
	"errors"
	"log"
	"simple-go-blog-webserver/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var ErrNoRows = errors.New("no rows found")
var ErrEmailIsOccupied = errors.New("email is occupied by another user")
var ErrUsernameIsOccupied = errors.New("username is occupied by another user")

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
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var usernameExists bool
	checkUsernameQuery := "SELECT EXISTS(SELECT 1 FROM blog_webserver_db.users WHERE username = ?)"
	err = tx.QueryRow(checkUsernameQuery, newUser.Username).Scan(&usernameExists)
	if err != nil {
		tx.Rollback()
		return err
	}
	if usernameExists {
		tx.Rollback()
		return ErrUsernameIsOccupied
	}

	var emailExists bool
	checkEmailQuery := "SELECT EXISTS(SELECT 1 FROM blog_webserver_db.users WHERE email = ?)"
	err = tx.QueryRow(checkEmailQuery, newUser.Email).Scan(emailExists)
	if err != nil {
		tx.Rollback()
		return err
	}
	if emailExists {
		tx.Rollback()
		return ErrEmailIsOccupied
	}

	_, err = tx.Exec("INSERT INTO blog_webserver_db.users (name, username, email) VALUES (?, ?, ?)", newUser.Name, newUser.Username, newUser.Email)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func GetUsers() ([]models.User, error) {
	rows, err := db.Query("SELECT id, name, username, email FROM blog_webserver_db.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Username,
			&user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func GetUserById(userId int) (models.User, error) {
	query := "SELECT * FROM blog_webserver_db.users WHERE id = ?"
	var user models.User
	err := db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNoRows
		}
		return user, err
	}
	return user, nil
}
