package handlers

import (
	"fmt"
	"musicplatform/config"
	"musicplatform/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {
	var songs []models.Song

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	genreName := c.Query("genre") // filter
	artistID := c.Query("artist_id")
	genreID := c.Query("genre_id")

	query := config.DB.Preload("Artist").Preload("Genre")

	if genreName != "" {
		query = query.Joins("JOIN genres ON genres.id = songs.genre_id").
			Where("genres.name = ?", genreName)
	}
	if artistID != "" {
		query = query.Where("artist_id = ?", artistID)
	}
	if genreID != "" {
		query = query.Where("genre_id = ?", genreID)
	}

	query.Offset(offset).Limit(limit).Find(&songs)

	c.JSON(http.StatusOK, songs)
}

func GetSongByID(c *gin.Context) {
	id := c.Param("id")
	var song models.Song

	if err := config.DB.Preload("Artist").Preload("Genre").First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func CreateSong(c *gin.Context) {
	var song models.Song
	if err := c.ShouldBindJSON(&song); err != nil {
		fmt.Println("Bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&song).Error; err != nil {
		fmt.Println("DB error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create song"})
		return
	}

	c.JSON(http.StatusCreated, song)
}

func UpdateSong(c *gin.Context) {
	id := c.Param("id")
	var song models.Song

	if err := config.DB.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&song)
	c.JSON(http.StatusOK, song)
}

func DeleteSong(c *gin.Context) {
	id := c.Param("id")
	var song models.Song

	if err := config.DB.First(&song, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	config.DB.Delete(&song)
	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}
