package post

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
)

func PostIndex(w http.ResponseWriter, r *http.Request) {
	posts, err := FindAll()
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(posts)
	fmt.Fprint(w, string(b))
}

func PostShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]

	post, err := FindById(postId)
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(post)
	fmt.Fprint(w, string(b))
}

func PostShowBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["slug"]

	post, err := FindBySlug(postId)
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	b, err := json.Marshal(post)
	fmt.Fprint(w, string(b))
}

func PostSave(w http.ResponseWriter, r *http.Request) {
	var _post Post

	// Get request body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}


	// Unmarshal
	if err := json.Unmarshal(body, &_post); err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	post, err := Save(_post)
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(post)
	fmt.Fprint(w, string(b))
}


func PostDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]

	err := Delete(postId)
	if err != nil {
		fmt.Println(err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}