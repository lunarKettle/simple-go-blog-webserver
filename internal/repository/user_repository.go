package repository

import (
	"database/sql"
	"errors"
	"simple-go-blog-webserver/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

var ErrNotFound = errors.New("no record found")
var ErrEmailIsOccupied = errors.New("email is occupied by another user")
var ErrUsernameIsOccupied = errors.New("username is occupied by another user")
var ErrFailToGetUsers = errors.New("failed to get users from database")

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(database Database) UserRepository {
	return UserRepository{db: database.connection}
}

func (ur *UserRepository) CreateUser(newUser models.User) error {
	tx, err := ur.db.Begin()
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

func (ur *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := ur.db.Query("SELECT id, name, username, email FROM blog_webserver_db.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, err
}

func (ur *UserRepository) GetUserById(userId int) (models.User, error) {
	query := "SELECT * FROM blog_webserver_db.users WHERE id = ?"
	var user models.User
	err := ur.db.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrNotFound
		}
		return user, err
	}
	return user, nil
}
