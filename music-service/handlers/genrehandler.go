package handlers

import (
	"net/http"
	"strconv"

	"musicplatform/music-service/config"
	"musicplatform/music-service/models"

	"github.com/gin-gonic/gin"
)

func GetGenres(c *gin.Context) {
	var genres []models.Genre
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	config.DB.Offset(offset).Limit(limit).Find(&genres)
	c.JSON(http.StatusOK, genres)
}

func GetGenreByID(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre

	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	c.JSON(http.StatusOK, genre)
}

func CreateGenre(c *gin.Context) {
	var genre models.Genre

	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&genre).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create genre"})
		return
	}

	c.JSON(http.StatusCreated, genre)
}

func UpdateGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre

	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&genre)
	c.JSON(http.StatusOK, genre)
}

func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	var genre models.Genre

	if err := config.DB.First(&genre, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Genre not found"})
		return
	}

	config.DB.Delete(&genre)
	c.JSON(http.StatusOK, gin.H{"message": "Genre deleted"})
}
