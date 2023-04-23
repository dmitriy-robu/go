package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Boxes []Box

type Box struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primaryKey" faker:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"unique"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	AltImage  string    `json:"alt_image"`
	Price     uint      `json:"price"`
	Active    bool      `json:"active" gorm:"default:0" faker:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BoxItems  BoxItems       `gorm:"foreignKey:BoxID"`
}
