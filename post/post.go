package post

import (
	"time"
)

type Post struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Slug    string    `json:"slug"`
	Desc    string    `json:"desc"`
	Content string    `json:"content"`
	Date    time.Time `json:"date"`
}

type Posts []Post
