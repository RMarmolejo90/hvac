package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rmarmolejo90/hvac/internal/app/initializers"
	"github.com/rmarmolejo90/hvac/internal/app/routes"
	"github.com/rmarmolejo90/hvac/internal/postgresDB"
)

func main() {
	// Initialize the application components
	postgresDB.ConnectDB()
	handlers := initializers.InitHandlers()

	// Set up the Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, handlers)

	// Start the server
	r.Run() // listen and serve on 0.0.0.0:8080
}
