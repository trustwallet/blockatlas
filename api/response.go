package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiError struct {
	*gin.Context  `json:"-"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

var EmptyResponse = map[string]interface{}{}

func RenderSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func RenderError(c *gin.Context, code int, msg string) {
	err := &apiError{
		StatusCode:    code,
		StatusMessage: msg,
		Context:       c,
	}
	err.Render()
}

func ErrorResponse(c *gin.Context) *apiError {
	return &apiError{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: "Internal server error",
		Context:       c,
	}
}

func (e *apiError) Code(code int) *apiError {
	e.StatusCode = code
	return e
}

func (e *apiError) Message(msg string) *apiError {
	e.StatusMessage = msg
	return e
}

func (e *apiError) Params(code int, msg string) *apiError {
	e.StatusCode = code
	e.StatusMessage = msg
	return e
}

func (e *apiError) Render() {
	e.AbortWithStatusJSON(e.StatusCode, e.StatusMessage)
}
