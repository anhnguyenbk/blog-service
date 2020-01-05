package main

import (
	"net/http"

	"github.com/akrylysov/algnhsa"
	"github.com/anhnguyenbk/blog-service/internal/post"
	"github.com/gorilla/mux"
)

// access control and  CORS middleware
func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

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
