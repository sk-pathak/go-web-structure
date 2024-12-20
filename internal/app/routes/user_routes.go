package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sk-pathak/go-structure/internal/app/handler"
)

// RegisterUserRoutes registers routes for user-related operations
func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)
}
