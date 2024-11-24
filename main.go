package main

import (
	"github.com/gin-gonic/gin"
	"mini-social-media/routes"
)

func main() {
	router := gin.Default()

	// Initialize routes
	routes.RegisterPostRoutes(router)

	// Start server
	router.Run(":8080")
}
