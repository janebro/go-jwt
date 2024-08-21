package main

import (
	"go-jwt/controllers"
	"go-jwt/database"
	"go-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	database.Connect("host=localhost user=go-jwt password=1234 dbname=go-jwt port=5432 sslmode=disable")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
