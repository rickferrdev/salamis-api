package response

import (
	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type StandardResponse struct {
	Status int    `json:"status"`
	Data   any    `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Response struct{}

func NewResponse() ports.Response {
	return &Response{}
}

func (u *Response) JSON(c *gin.Context, status int, data any) {
	c.JSON(200)
}

func format(err any) string {
	switch vl := err.(type) {
	case error:
		return vl.Error()
	case string:
		return vl
	default:
		return "error processing a response"
	}
}
