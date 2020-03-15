package post

import (
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/helper"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var _comment Comment

	postId := helper.ParsePathParam(r, "postId")
	err := helper.ParseJSONBody(r, &_comment)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}

	comment, err := SaveComment(postId, _comment)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, comment)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	postId := helper.ParsePathParam(r, "postId")

	comments, err := PostGetComments(postId)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
	helper.ResponseJSON(w, comments)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	postId := helper.ParsePathParam(r, "postId")
	commentId := helper.ParsePathParam(r, "commentId")

	err := PostDeleteComment(postId, commentId)
	if err != nil {
		helper.ResponseError(w, err)
		return
	}
}
