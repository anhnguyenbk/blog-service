package shared

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
