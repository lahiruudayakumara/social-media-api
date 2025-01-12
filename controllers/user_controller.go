package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"social-media-api/config"
	"social-media-api/models"
)

// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Preload("Posts").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Get a user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.Preload("Posts").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update a user profile
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete a user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Find the user by ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
