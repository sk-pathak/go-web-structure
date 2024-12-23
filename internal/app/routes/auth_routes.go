package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/sk-pathak/go-structure/internal/app/handler"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	r.GET("/auth/begin", authHandler.BeginAuth)
	r.GET("/auth/complete", authHandler.CompleteAuth)
	r.GET("/auth/logout", authHandler.Logout)
	r.GET("/auth/session", authHandler.GetUserSession)
}
