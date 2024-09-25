package main

import (
	"crud-app/config"
	"crud-app/database"
	"crud-app/handler"
	"crud-app/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	serveApplication()
}

func serveApplication() {

	db := database.Connect()
	// Initialize Gin
	router := gin.Default()
	router.Use(middleware.ContextMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Group routes
	authRoutes := router.Group("/v1/public")
	{
		handler.SetupAuthRoutes(authRoutes, db)
	}

	// Set up routes
	api := router.Group("/v1", middleware.AuthMiddleware())
	{
		handler.SetupProfile(api, db)
		// You can add more route setups for other models if needed
	}

	router.Run(":8001")
	fmt.Println("Server running on port 8000")
}
