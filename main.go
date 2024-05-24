package main

import (
	"github.com/ApesJs/test-be/initializers"
	"github.com/ApesJs/test-be/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	route := gin.Default()

	// Konfigurasi CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5174"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	// Gunakan middleware CORS
	route.Use(cors.New(config))

	routes.PostRoutes(route)

	if err := route.Run(os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
