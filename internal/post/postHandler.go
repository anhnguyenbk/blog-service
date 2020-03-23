package post

import (
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/util/requestutils"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := FindAll()
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, posts)
}

func GetPostsByCateSlugHandler(w http.ResponseWriter, r *http.Request) {
	category := requestutils.ParsePathParam(r, "slug")

	posts, err := FindByCategory(category)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, posts)
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId := requestutils.ParsePathParam(r, "id")

	post, err := FindById(postId)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, post)
}

func GetPostBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := requestutils.ParsePathParam(r, "slug")

	post, err := FindBySlug(slug)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, post)
}

func SavePostHandler(w http.ResponseWriter, r *http.Request) {
	var _post = Post{}
	err := requestutils.ParseJSONBody(r, &_post)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}

	post, err := Save(_post)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}

	responseutils.ResponseJSON(w, post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	postId := requestutils.ParsePathParam(r, "id")

	err := Delete(postId)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
}
