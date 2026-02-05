package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type AuthHandler struct {
	response ports.Response
	service  ports.AuthService
}

func NewAuthHandler(service ports.AuthService, response ports.Response) *AuthHandler {
	return &AuthHandler{
		service:  service,
		response: response,
	}
}

func (u *AuthHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		500*time.Millisecond,
	)
	defer cancel()

	var body RequestUserLoginDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ports.ErrInvalidInput.Error(),
			"details": err.Error(),
		})
		return
	}

	output, err := u.service.Login(ctx, ports.AuthInput{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		u.response.Error(ctx, c.Writer.(http.ResponseWriter), err)
		return
	}

	u.response.Success(ctx, c.Writer.(http.ResponseWriter), 200, ResponseUserLoginDTO{
		Nickname: output.User.Nickname,
		Token:    output.Token,
	})
}

func (u *AuthHandler) Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(
		c.Request.Context(),
		500*time.Millisecond,
	)
	defer cancel()

	var body RequestUserRegisterDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   ports.ErrInvalidInput.Error(),
			"details": err.Error(),
		})
		return
	}

	output, err := u.service.Register(ctx, ports.AuthInput{
		Nickname: body.Nickname,
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		u.response.Error(ctx, c.Writer.(http.ResponseWriter), err)
		return
	}

	u.response.Success(ctx, c.Writer.(http.ResponseWriter), 200, ResponseUserRegisterDTO{
		Nickname: output.User.Nickname,
	})
}
