package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/sk-pathak/go-structure/internal/app/handler"
	"github.com/sk-pathak/go-structure/internal/middlewares"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler, store sessions.Store) {
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users", userHandler.GetUsers)

	protectedGroup := r.Group("/users")
	protectedGroup.Use(middlewares.AuthMiddleware(store))
	{
		protectedGroup.GET("/:id", userHandler.GetUser)
	}
}
