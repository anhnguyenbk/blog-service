package post // import "github.com/anhnguyenbk/blog-service/internal/post"

import "time"

type Category struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type Categories []Category

type Post struct {
	Id         string     `json:"id"`
	Title      string     `json:"title"`
	Slug       string     `json:"slug"`
	Desc       string     `json:"desc"`
	Content    string     `json:"content"`
	Status     string     `json:"status"`
	Categories Categories `json:"categories"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

type Posts []Post
