package models

import "time"

type Post struct {
	Id        int
	Text      string
	UserId    int
	Date      time.Time
	IsChanged bool
}
