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

func (pr *PostRepository) GetPostById(postId int) (models.Post, error) {
	query := "SELECT * FROM blog_webserver_db.posts WHERE id = ?"
	var post models.Post
	err := pr.db.QueryRow(query, postId).Scan(&post.Id, &post.Text, &post.UserId, &post.Date, &post.IsChanged)
	if err != nil {
		if err == sql.ErrNoRows {
			return post, ErrNotFound
		}
		return post, err
	}
	return post, nil
}
