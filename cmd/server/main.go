package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	config "github.com/sk-pathak/go-structure/configs"
	"github.com/sk-pathak/go-structure/internal/app/handler"
	"github.com/sk-pathak/go-structure/internal/app/repository"
	"github.com/sk-pathak/go-structure/internal/app/routes"
	"github.com/sk-pathak/go-structure/internal/app/service"
	"github.com/sk-pathak/go-structure/internal/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	connStr := "user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" host=" + cfg.DBHost +
		" port=" + cfg.DBPort +
		" dbname=" + cfg.DBName +
		" sslmode=disable"

	// Establish connection to PostgreSQL using pgx
	pgConn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Ensure the connection is open and working
	if err := pgConn.Ping(context.Background()); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	// Initialize the db.Queries instance using the pgx connection
	queries := db.New(pgConn)

	// Create repositories using the generated queries
	userRepo := repository.NewUserRepository(queries)

	// Create services
	userService := service.NewUserService(userRepo)

	// Create handlers
	userHandler := handler.NewUserHandler(userService)

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterUserRoutes(r, userHandler)

	// Start the server using the port defined in the configuration
	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
