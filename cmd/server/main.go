package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sk-pathak/go-structure/configs"
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

    connStr := cfg.DBDriver + "://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName

    dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
    defer dbPool.Close()

	if err := dbPool.Ping(context.Background()); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	queries := db.New(dbPool)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	routes.RegisterUserRoutes(r, userHandler)

	log.Printf("Server is running on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}
