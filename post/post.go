package post

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Desc      string    `json:"desc"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Posts []Post
