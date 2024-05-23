package main

import (
	"github.com/ApesJs/test-be/initializers"
	"github.com/ApesJs/test-be/routes"
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

	routes.PostRoutes(route)

	if err := route.Run(os.Getenv("PORT")); err != nil {
		panic(err)
	}
}
