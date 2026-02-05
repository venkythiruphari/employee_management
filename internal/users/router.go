package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup, handler *Handler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", handler.Register)
		userRoutes.POST("/login", handler.Login)
	}
}
