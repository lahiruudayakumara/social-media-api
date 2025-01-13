package routes

import (
	"github.com/gin-gonic/gin"
	"social-media-api/controllers"
)

func SetupAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
