package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

// User полная таблица в mysql
type User struct {
	ID              *uint64 `gorm:"primaryKey"`
	UUID            string  `gorm:"unique"`
	Name            *string
	AvatarURL       *string
	Email           *string
	EmailVerifiedAt *time.Time
	Password        *string
	SteamTradeURL   *string
	Experience      *uint64
	Active          bool
	IsBot           bool
	RememberToken   *string
	DeletedAt       *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	gorm.Model
}

type UserSteamInfo struct {
	SteamID   *string
	AvatarURL *string
	Name      *string
}

type UserSteamProfile struct {
	SteamID   *string `json:"steamid"`
	AvatarURL *string `json:"avatar"`
	Name      *string `json:"personaname"`
}

type UserWithBalance struct {
	User        User
	UserBalance UserBalance
}

type UserAuthSteam struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	UserID       *uint64            `bson:"user_id"`
	SteamID      *string            `bson:"steam_id"`
	Token        *string            `bson:"token"`
	RefreshToken *string            `bson:"refresh_token"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}
