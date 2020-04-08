package model

const (
	Default ErrorCode = iota
	InvalidQuery
	RequestedDataNotFound
	InternalFail
)

type (
	ErrorResponse struct {
		Error ErrorDetails `json:"error"`
	}
	ErrorDetails struct {
		Message string    `json:"message"`
		Code    ErrorCode `json:"code"`
	}

	ErrorCode int
)

func CreateErrorResponse(code ErrorCode, err error) ErrorResponse {
	var message string
	if err != nil {
		message = err.Error()
	}
	return ErrorResponse{Error: ErrorDetails{
		Message: message,
		Code:    code,
	}}
}
