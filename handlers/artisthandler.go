package handlers

import (
	"musicplatform/config"
	"musicplatform/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArtists(c *gin.Context) {
	var artists []models.Artist
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	config.DB.Offset(offset).Limit(limit).Find(&artists)
	c.JSON(http.StatusOK, artists)
}

func GetArtistByID(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := config.DB.First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}

	c.JSON(http.StatusOK, artist)
}

func CreateArtist(c *gin.Context) {
	var artist models.Artist

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&artist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create artist"})
		return
	}

	c.JSON(http.StatusCreated, artist)
}

func UpdateArtist(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := config.DB.First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}

	if err := c.ShouldBindJSON(&artist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&artist)
	c.JSON(http.StatusOK, artist)
}

func DeleteArtist(c *gin.Context) {
	id := c.Param("id")
	var artist models.Artist

	if err := config.DB.First(&artist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Artist not found"})
		return
	}

	config.DB.Delete(&artist)
	c.JSON(http.StatusOK, gin.H{"message": "Artist deleted"})
}
