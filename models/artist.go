package models

import (
	// "encoding/json"
	"time"
)

type Artist struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"-" json:"-"`
	UpdatedAt time.Time  `gorm:"-" json:"-"`
	DeletedAt *time.Time `gorm:"-" json:"-"`
	Name      string     `json:"name"`
}
