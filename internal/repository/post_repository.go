package repository

import (
	"database/sql"
	"simple-go-blog-webserver/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(database Database) PostRepository {
	return PostRepository{db: database.connection}
}

func (pr *PostRepository) AddPost(newPost models.Post) error {
	query := "INSERT INTO blog_webserver_db.posts (text, userId, date, isChanged) VALUES (?, ?, ?, ?)"
	_, err := pr.db.Exec(query, newPost.Text, newPost.UserId, newPost.Date, newPost.IsChanged)
	return err
}
