package shared

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ResponseError struct {
	Error CodeError `json:"error"`
}

func ToResponseError(err error) ResponseError {
	var codeError = CodeError{500, err.Error()}
	return ResponseError{codeError}
}
