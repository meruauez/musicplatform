package models

import (
	"encoding/json"
	"time"
)

type Genre struct {
	ID        uint       `gorm:"primaryKey" json:"id"` // ID — остается видимым
	CreatedAt time.Time  `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	UpdatedAt time.Time  `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	DeletedAt *time.Time `gorm:"-" json:"-"`           // Исключаем из JSON и миграций
	Name      string     `json:"name"`                 // Только имя будет отображаться в ответе
	// Связь с песнями, если не нужно показывать, исключаем
	// Songs     []Song `gorm:"foreignKey:GenreID" json:"-"`
}

// MarshalJSON позволяет контролировать, как будет сериализоваться объект в JSON
func (g *Genre) MarshalJSON() ([]byte, error) {
	type Alias Genre
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(g),
	})
}
