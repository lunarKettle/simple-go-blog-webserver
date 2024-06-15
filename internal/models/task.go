package models

import "time"

type Task struct {
	Text        string
	IsCompleted bool
	Date        time.Time
}
