package routes

import (
	"github.com/ApesJs/test-be/controllers"
	"github.com/gin-gonic/gin"
)

func PostRoutes(route *gin.Engine) {
	route.POST("/posts", controllers.CreatePost)
	route.GET("/posts", controllers.FindAllPosts)
	route.GET("/posts/page/:page", controllers.FindAllPosts)
	route.GET("/posts/:id", controllers.FindByIDPost)
	route.PUT("/posts/:id", controllers.UpdatePost)
	route.DELETE("/posts/:id", controllers.DeletePost)
	route.PUT("/posts/trashed/:id", controllers.TrashedPost)
}
