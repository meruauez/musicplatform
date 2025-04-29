package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres" // <--- меняем
	"gorm.io/gorm"

	"musicplatform/config"
	"musicplatform/models"
)

func setupTestDB() {
	dsn := "host=localhost user=postgres password=postgre17 dbname=musicplatform_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Делаем миграции для нужных таблиц
	db.AutoMigrate(&models.Song{}, &models.Artist{}, &models.Genre{})

	// Присваиваем глобальной переменной
	config.DB = db
}

func TestGetSongs(t *testing.T) {
	gin.SetMode(gin.TestMode)

	setupTestDB()

	// Наполним тестовыми данными
	config.DB.Create(&models.Song{
		Title: "Test Song",
	})

	router := gin.Default()
	router.GET("/songs", GetSongs)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetSongByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Сначала создаём тестового артиста
	artist := models.Artist{Name: "Test Artist"}
	config.DB.Create(&artist)

	// Потом создаём тестовый жанр
	genre := models.Genre{Name: "Test Genre"}
	config.DB.Create(&genre)

	// Потом создаём песню с валидным artist_id и genre_id
	song := models.Song{
		Title:    "Song 1",
		ArtistID: artist.ID,
		GenreID:  genre.ID,
	}
	config.DB.Create(&song)

	// Теперь настраиваем роутер и тестируем
	router := gin.Default()
	router.GET("/songs/:id", GetSongByID)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs/"+strconv.Itoa(int(song.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestCreateSong(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	artist := models.Artist{Name: "Test Artist"}
	config.DB.Create(&artist)

	genre := models.Genre{Name: "Test Genre"}
	config.DB.Create(&genre)

	song := models.Song{
		Title:    "New Song",
		ArtistID: artist.ID,
		GenreID:  genre.ID,
	}
	payload, _ := json.Marshal(song)

	router := gin.Default()
	router.POST("/songs", CreateSong)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/songs", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestUpdateSong(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Добавляем тестовые Artist и Genre
	artist := models.Artist{Name: "Test Artist"}
	config.DB.Create(&artist)

	genre := models.Genre{Name: "Test Genre"}
	config.DB.Create(&genre)

	// Добавляем песню с правильными связями
	song := models.Song{
		Title:    "Song 1",
		ArtistID: artist.ID,
		GenreID:  genre.ID,
	}
	config.DB.Create(&song)

	// Меняем заголовок песни
	song.Title = "Updated Song"
	payload, _ := json.Marshal(song)

	router := gin.Default()
	router.PUT("/songs/:id", UpdateSong)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/songs/"+strconv.Itoa(int(song.ID)), bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeleteSong(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Добавляем тестовые Artist и Genre
	artist := models.Artist{Name: "Test Artist"}
	config.DB.Create(&artist)

	genre := models.Genre{Name: "Test Genre"}
	config.DB.Create(&genre)

	// Добавляем песню с правильными связями
	song := models.Song{
		Title:    "Song 1",
		ArtistID: artist.ID,
		GenreID:  genre.ID,
	}
	config.DB.Create(&song)

	router := gin.Default()
	router.DELETE("/songs/:id", DeleteSong)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/songs/"+strconv.Itoa(int(song.ID)), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetSongsByArtist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Добавляем тестового артиста и жанр
	artist := models.Artist{Name: "Artist 1"}
	config.DB.Create(&artist)

	genre := models.Genre{Name: "Genre 1"}
	config.DB.Create(&genre)

	// Добавляем песню с правильными связями
	song := models.Song{
		Title:    "Song 1",
		ArtistID: artist.ID,
		GenreID:  genre.ID,
	}
	config.DB.Create(&song)

	router := gin.Default()
	router.GET("/artists/:id/songs", GetSongsByArtist)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/artists/"+strconv.Itoa(int(artist.ID))+"/songs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetSongsByGenre(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Добавляем тестовый жанр
	genre := models.Genre{Name: "Genre 1"}
	config.DB.Create(&genre)

	// Добавляем песню с правильным жанром
	song := models.Song{
		Title:   "Song 1",
		GenreID: genre.ID,
	}
	config.DB.Create(&song)

	router := gin.Default()
	router.GET("/genres/:id/songs", GetSongsByGenre)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/genres/"+strconv.Itoa(int(genre.ID))+"/songs", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestSearchSongs(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Добавляем артиста и жанр для предотвращения ошибок внешнего ключа
	artist := models.Artist{Name: "Artist 1"}
	config.DB.Create(&artist)

	genre := models.Genre{Name: "Genre 1"}
	config.DB.Create(&genre)

	// Создаем песни с правильными artist_id и genre_id
	config.DB.Create(&models.Song{Title: "Test Song", ArtistID: artist.ID, GenreID: genre.ID})
	config.DB.Create(&models.Song{Title: "Another Song", ArtistID: artist.ID, GenreID: genre.ID})

	router := gin.Default()
	router.GET("/songs/search", SearchSongs)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/songs/search?q=Test", nil)
	router.ServeHTTP(w, req)

	// Проверяем, что код ответа 200
	assert.Equal(t, 200, w.Code)

	// Проверяем, что в ответе есть песня с названием "Test Song"
	var songs []models.Song
	err := json.Unmarshal(w.Body.Bytes(), &songs)
	assert.NoError(t, err)
	assert.Len(t, songs, 1)
	assert.Equal(t, "Test Song", songs[0].Title)
}
