package models

type Song struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CreatedAt string  `gorm:"-" json:"-"`
	UpdatedAt string  `gorm:"-" json:"-"`
	DeletedAt *string `gorm:"-" json:"-"`
	Title     string  `json:"title"`
	ArtistID  uint    `json:"artist_id"`
	Artist    Artist  `gorm:"foreignKey:ArtistID" json:"-"`
	GenreID   uint    `json:"genre_id"`
	Genre     Genre   `gorm:"foreignKey:GenreID" json:"-"`
}
