package models

import (
	"encoding/json"
	"time"
)

type Artist struct {
	ID        uint       `gorm:"primaryKey" json:"id"` // ID — остается видимым
	CreatedAt time.Time  `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	UpdatedAt time.Time  `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	DeletedAt *time.Time `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	Name      string     `json:"name"`                 // Только имя будет отображаться в ответе
	// Связь с песнями, если не нужно показывать, исключаем
	// Songs     []Song `gorm:"foreignKey:ArtistID" json:"-"`
}

// MarshalJSON позволяет контролировать, как будет сериализоваться объект в JSON
func (a *Artist) MarshalJSON() ([]byte, error) {
	type Alias Artist
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	})
}
