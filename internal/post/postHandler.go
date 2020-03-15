package post

import (
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/helper"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := FindAll()
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, posts)
}

func GetPostsByCateSlugHandler(w http.ResponseWriter, r *http.Request) {
	category := helper.ParsePathParam(r, "slug")

	posts, err := FindByCategory(category)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId := helper.ParsePathParam(r, "id")

	post, err := FindById(postId)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, post)
}

func GetPostBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := helper.ParsePathParam(r, "slug")

	post, err := FindBySlug(slug)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, post)
}

func SavePostHandler(w http.ResponseWriter, r *http.Request) {
	var _post = Post{}
	err := helper.ParseJSONBody(r, &_post)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	post, err := Save(_post)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	helper.ResponseJSON(w, post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	postId := helper.ParsePathParam(r, "id")

	err := Delete(postId)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
}
