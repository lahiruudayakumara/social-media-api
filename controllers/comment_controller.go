package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-api/config"
	"social-media-api/models"
	"strconv" // Import strconv for conversion
)

// Get comments for a specific post
func GetCommentsByPostID(c *gin.Context) {
	var comments []models.Comment
	postID := c.Param("id") // The id is a strings

	// Convert postID to uint
	postIDUint, err := strconv.ParseUint(postID, 10, 32) // Convert to uint32 (for PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Query the database for comments with the id
	if err := config.DB.Where("post_id = ?", uint(postIDUint)).Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// Create a new comment for a post
func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get id from the URL and convert it to uint
	postID := c.Param("id")
	postIDUint, err := strconv.ParseUint(postID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comment.PostID = uint(postIDUint) // Set the PostID field to the converted uint

	// Create the new comment in the database
	if err := config.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

// Delete a comment by ID
func DeleteComment(c *gin.Context) {
	commentID := c.Param("id") // Get the comment_id from the route parameters

	// Convert commentID to uint
	commentIDUint, err := strconv.ParseUint(commentID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	// Find and delete the comment
	var comment models.Comment
	if err := config.DB.Where("id = ?", uint(commentIDUint)).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	if err := config.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
