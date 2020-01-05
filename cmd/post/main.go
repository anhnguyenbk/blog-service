package main

import (
	"github.com/akrylysov/algnhsa"
	"github.com/anhnguyenbk/blog-service/internal/post"
	"github.com/gorilla/mux"
	// "net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", post.PostIndex).Methods("GET")
	router.HandleFunc("/", post.PostSave).Methods("POST")
	router.HandleFunc("/slug/{slug}", post.PostShowBySlug).Methods("GET")
	router.HandleFunc("/{id}", post.PostShow).Methods("GET")
	router.HandleFunc("/{id}", post.PostSave).Methods("POST")
	router.HandleFunc("/{id}", post.PostDelete).Methods("DELETE")

	// Local http
	// http.ListenAndServe(":3000", nil)

	// Lambda
	algnhsa.ListenAndServe(router, nil)
}
