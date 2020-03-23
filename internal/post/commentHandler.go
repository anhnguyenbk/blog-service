package post

import (
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/util/requestutils"
	"github.com/anhnguyenbk/blog-service/internal/util/responseutils"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	var _comment Comment

	postId := requestutils.ParsePathParam(r, "postId")
	err := requestutils.ParseJSONBody(r, &_comment)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}

	comment, err := SaveComment(postId, _comment)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, comment)
}

func GetCommentsHandler(w http.ResponseWriter, r *http.Request) {
	postId := requestutils.ParsePathParam(r, "postId")

	comments, err := PostGetComments(postId)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
	responseutils.ResponseJSON(w, comments)
}

func DeleteCommentHandler(w http.ResponseWriter, r *http.Request) {
	postId := requestutils.ParsePathParam(r, "postId")
	commentId := requestutils.ParsePathParam(r, "commentId")

	err := PostDeleteComment(postId, commentId)
	if err != nil {
		responseutils.ResponseError(w, err)
		return
	}
}
