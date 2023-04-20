package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Boxes []Box

type Box struct {
	ID       uint      `json:"id" gorm:"primaryKey" faker:"-"`
	UUID     uuid.UUID `json:"uuid" gorm:"unique"`
	Title    string    `json:"title" `
	Image    string    `json:"image"`
	AltImage string    `json:"alt_image"`
	Price    int       `json:"price"`
	Active   int       `json:"active" gorm:"default:0" faker:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	BoxItem   []BoxItem      `gorm:"foreignKey:BoxID"`
}
