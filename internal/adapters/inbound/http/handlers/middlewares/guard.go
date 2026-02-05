package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type GuardMiddleware struct {
	tokenizer ports.Tokenizer
	response  ports.Response
}

func NewGuardMiddleware(tokenizer ports.Tokenizer, response ports.Response) *GuardMiddleware {
	return &GuardMiddleware{
		tokenizer: tokenizer,
		response:  response,
	}
}

func (u *GuardMiddleware) Protect() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.Split(header, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			u.response.Error(context.TODO(), c.Writer.(http.ResponseWriter), ports.ErrInvalidFormatAuthorization)
			c.Abort()
			return
		}

		claims, err := u.tokenizer.Verify(parts[1])
		if err != nil {
			if errors.Is(err, ports.ErrInvalidToken) {
				u.response.Error(context.TODO(), c.Writer.(http.ResponseWriter), err)
				c.Abort()
				return
			}

			u.response.Error(context.TODO(), c.Writer.(http.ResponseWriter), ports.ErrInternalServer)
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
