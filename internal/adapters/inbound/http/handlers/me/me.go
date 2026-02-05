package me

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rickferrdev/salamis-api/internal/core/ports"
)

type MeHandler struct {
	service ports.AuthService
}

func NewMeHandler(service ports.AuthService) *MeHandler {
	return &MeHandler{
		service: service,
	}
}

func (u *MeHandler) Me(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": ports.ErrUnauthorized.Error()})
		return
	}

	ctx.JSON(200, gin.H{"user_id": userID})
}
