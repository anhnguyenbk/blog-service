package helper // import "github.com/anhnguyenbk/blog-service/internal/helper"

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anhnguyenbk/blog-service/internal/shared"
	"github.com/gorilla/mux"
)

func ParseJSONBody(r *http.Request, dst interface{}) error {
	// Get request body
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		return err
	}

	return nil
}

func ParsePathParam(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

func ResponseError(w http.ResponseWriter, err error) {
	// Log the error
	fmt.Println(err.Error())

	_json, err := json.Marshal(shared.ToResponseError(err))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, string(_json))
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	_json, err := json.Marshal(data)
	if err != nil {
		ResponseError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(_json))
}
