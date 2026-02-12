package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type authHandler struct {
	service  ports.AuthService
	response ports.Response
}

func NewAuthHandler(service ports.AuthService, router gin.RouterGroup) *authHandler {
	handler := authHandler{
		service: service,
	}

	group := router.Group("/auth")
	group.GET("/login", handler.Login)
	group.GET("/register", handler.Register)

	return &handler
}

func (u *authHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 25*time.Second)
	defer cancel()

	var body RequestLoginDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "poorly formatted request body"})
		return
	}

	output, err := u.service.Login(ctx, ports.AuthInput{Password: body.Password, Username: body.Username})
	if err != nil {
		if errors.Is(err, ports.ErrUserNotFound) {
			c.AbortWithStatusJSON(404, gin.H{"error": ports.ErrUserAlreadyExists.Error()})
			return
		}

		if errors.Is(err, ports.ErrInvalidCredentials) {
			c.AbortWithStatusJSON(401, gin.H{"error": ports.ErrInvalidCredentials.Error()})
			return
		}

		if errors.Is(ctx.Err(), context.DeadlineExceeded) || errors.Is(ctx.Err(), context.Canceled) {
			c.(c.Writer.(http.ResponseWriter), 408, gin.H{"erorr": ports.ErrTimeout})
			return
		}

		u.response.JSON(c.Writer.(http.ResponseWriter), 500, gin.H{"error": ports.ErrInternalServerError})
		return
	}

	u.response.JSON(c.Writer.(http.ResponseWriter), 200, ResponseLoginDTO{Token: output.Token})
}

func (u *authHandler) Register(c *gin.Context) {
	// TODO: add error
}
