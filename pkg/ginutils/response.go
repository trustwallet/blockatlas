package ginutils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiError struct {
	*gin.Context  `json:"-"`
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
}

func RenderSuccess(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func RenderError(c *gin.Context, code int, msg string) {
	err := &ApiError{
		StatusCode:    code,
		StatusMessage: msg,
		Context:       c,
	}
	err.Render()
}

func ErrorResponse(c *gin.Context) *ApiError {
	return &ApiError{
		StatusCode:    http.StatusInternalServerError,
		StatusMessage: "Internal server error",
		Context:       c,
	}
}

func (e *ApiError) Code(code int) *ApiError {
	e.StatusCode = code
	return e
}

func (e *ApiError) Message(msg string) *ApiError {
	e.StatusMessage = msg
	return e
}

func (e *ApiError) Params(code int, msg string) *ApiError {
	e.StatusCode = code
	e.StatusMessage = msg
	return e
}

func (e *ApiError) Render() {
	var msg map[string]interface{}
	err := json.Unmarshal([]byte(e.StatusMessage), &msg)
	if err == nil {
		e.AbortWithStatusJSON(e.StatusCode, msg)
		return
	}
	e.AbortWithStatusJSON(e.StatusCode, e.StatusMessage)
}
