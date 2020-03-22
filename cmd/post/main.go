package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/anhnguyenbk/blog-service/internal/helper"
	"github.com/anhnguyenbk/blog-service/internal/post"
	"github.com/anhnguyenbk/blog-service/internal/user"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(helper.CORSMiddleware, helper.LoggerMiddleware)

	// Users
	router.HandleFunc("/auth/token", user.AuthenticateHandler).Methods("POST", "OPTIONS")

	// CommentAPI
	router.HandleFunc("/posts/{postId}/comments", post.GetCommentsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/{postId}/comments", post.AddCommentHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts/{postId}/comments/{commentId}", post.DeleteCommentHandler).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/posts", post.GetPostsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts", post.SavePostHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts/slug/{slug}", post.GetPostBySlugHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/categories/{slug}", post.GetPostsByCateSlugHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/{id}", post.GetPostHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/posts/{id}", post.SavePostHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/posts/{id}", post.DeletePostHandler).Methods("DELETE", "OPTIONS")

	// Local http
	//http.ListenAndServe(":8080", router)

	// AWS Lambda
	algnhsa.ListenAndServe(router, nil)
}
