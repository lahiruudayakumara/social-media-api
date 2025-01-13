package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-api/config"
	"social-media-api/models"
)

// Get all posts
func GetPosts(c *gin.Context) {
	var posts []models.Post
	if err := config.DB.Preload("Comments").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Create a new post
func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// Get a single post by ID
func GetPostByID(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	if err := config.DB.Preload("Comments").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// Delete a post by ID
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
