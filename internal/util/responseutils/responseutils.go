package responseutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
type ErrorResponse struct {
	Error StatusError `json:"error"`
}

func ToStatusError(err error) StatusError {
	return StatusError{500, err.Error()}
}

func ToErrorResponse(statusError StatusError) ErrorResponse {
	return ErrorResponse{statusError}
}

func ResponseError(w http.ResponseWriter, err error) {
	// Log the error
	fmt.Println(err.Error())
	_json, err := json.Marshal(ToErrorResponse(ToStatusError(err)))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, string(_json))
}

func ResponseErrorWithStatus(w http.ResponseWriter, status int, err error) {
	// Log the error
	fmt.Println(err.Error())
	_json, err := json.Marshal(ToErrorResponse(StatusError{status, err.Error()}))
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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
