package models

type Song struct {
	ID        uint    `gorm:"primaryKey" json:"-"`
	CreatedAt string  `gorm:"-" json:"-"`                   // Исключаем из JSON и миграций
	UpdatedAt string  `gorm:"-" json:"-"`                   // Исключаем из JSON и миграций
	DeletedAt *string `gorm:"-" json:"-"`                   // Исключаем из JSON и миграций
	Title     string  `json:"title"`                        // Только заголовок песни будет отображаться в ответе
	ArtistID  uint    `json:"artist_id"`                    // ID артиста
	Artist    Artist  `gorm:"foreignKey:ArtistID" json:"-"` // Исключаем информацию о самом артисте
	GenreID   uint    `json:"genre_id"`                     // ID жанра
	Genre     Genre   `gorm:"foreignKey:GenreID" json:"-"`  // Исключаем информацию о самом жанре
}
