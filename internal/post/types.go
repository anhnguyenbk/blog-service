package post // import "github.com/anhnguyenbk/blog-service/internal/post"

import "time"

type Category struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type Categories []Category

type Comment struct {
	Id        string    `json:"id"`
	User      string    `json:"user"`
	Content   string    `json:"content"`
	Reply     Comments  `json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Comments []Comment

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
	Comments   Comments   `json:"comments"`
}

type Posts []Post
