package routes

import (
	"github.com/gin-gonic/gin"
	"social-media-api/controllers"
	"social-media-api/middleware"
)

func SetupUserRoutes(router *gin.Engine) {
	user := router.Group("/users")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}
