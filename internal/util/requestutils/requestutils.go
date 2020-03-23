package requestutils

import (
	"encoding/json"
	"net/http"

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
