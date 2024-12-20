package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sk-pathak/go-structure/internal/app/service"
	"github.com/sk-pathak/go-structure/internal/db"
)

type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Extract context from Gin's context
	ctx := c.Request.Context()

	if err := h.service.CreateUser(ctx, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers handles GET /users
func (h *UserHandler) GetUsers(c *gin.Context) {
	// Extract context from Gin's context
	ctx := c.Request.Context()

	users, err := h.service.GetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
