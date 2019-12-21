package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/anhnguyenbk/blog-service/post"
	"github.com/gorilla/mux"
)

type errorHandler func(http.ResponseWriter, *http.Request) (int, error)

// Our appHandler type will now satisify http.Handler
func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if status, err := fn(w, r); err != nil {
		fmt.Println(err)

		switch status {
		// We can have cases as granular as we like, if we wanted to
		// return custom errors for specific status codes.
		case http.StatusNotFound:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, `{"code":404,"message":%q}`, "Not found: "+err.Error())
			return
		// case http.StatusInternalServerError:
		// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		default:
			// Catch any other errors we haven't explicitly handled
			// http.Error(w, ErrorResponse{status, err}, http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"code":500,"message":%q}`, "An unexpected error has occurred: "+err.Error())
			return
		}
	}
}

func PostIndex(w http.ResponseWriter, r *http.Request) (int, error) {
	posts, err := post.FindAll()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	json.NewEncoder(w).Encode(posts)
	return http.StatusOK, nil
}

func PostSave(w http.ResponseWriter, r *http.Request) (int, error) {
	var _post post.Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if err := json.Unmarshal(body, &_post); err != nil {
		return http.StatusInternalServerError, err
	}

	post, err := post.Save(_post)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	json.NewEncoder(w).Encode(post)
	return http.StatusOK, nil
}

func PostShow(w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	postID := vars["postId"]
	post, err := post.FindById(postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	json.NewEncoder(w).Encode(post)
	return http.StatusOK, nil
}
