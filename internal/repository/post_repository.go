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

func (pr *PostRepository) GetPostsByUserId(userId int) ([]models.Post, error) {
	query := "SELECT * FROM blog_webserver_db.posts WHERE userId = ?"
	rows, err := pr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Text, &post.UserId, &post.Date, &post.IsChanged)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return posts, err
	}
	return posts, err
}
