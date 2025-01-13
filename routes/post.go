package routes

import (
	"github.com/gin-gonic/gin"
	"social-media-api/controllers"
	"social-media-api/middleware"
)

func SetupPostRoutes(router *gin.Engine) {
	post := router.Group("/posts")
	post.Use(middleware.AuthMiddleware())
	{
		post.GET("/", controllers.GetPosts)
		post.POST("/", controllers.CreatePost)
		post.GET("/:id", controllers.GetPostByID)
		post.DELETE("/:id", controllers.DeletePost)

		// Define your routes
		post.GET("/:id/comments", controllers.GetCommentsByPostID) // Get comments for a specific post
		post.POST("/:id/comments", controllers.CreateComment)      // Create a new comment for a post 		// Get all comments across all posts
		post.DELETE("/comments/:id", controllers.DeleteComment)    // Delete a comment by its ID

	}
}
