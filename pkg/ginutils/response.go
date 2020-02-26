package ginutils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	*gin.Context `json:"-"`
	Error        ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Message string `json:"message"`
}

func RenderSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func RenderError(c *gin.Context, code int, msg string) {
	details := ErrorDetails{
		Message: msg,
	}
	err := &ApiError{
		Error:   details,
		Context: c,
	}
	err.Render(code)
}

func ErrorResponse(c *gin.Context) *ApiError {
	details := ErrorDetails{
		Message: "Internal server error",
	}
	return &ApiError{
		Error:   details,
		Context: c,
	}
}

func (e *ApiError) Message(msg string) *ApiError {
	e.Error.Message = msg
	return e
}

func (e *ApiError) Params(msg string) *ApiError {
	e.Error.Message = msg
	return e
}

func (e *ApiError) Render(httpCode int) {
	e.AbortWithStatusJSON(httpCode, &e)
}
