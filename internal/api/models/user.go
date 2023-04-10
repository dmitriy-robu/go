package models

import (
	"database/sql"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"time"
)

// User полная таблица в mysql
type User struct {
	gorm.Model
	ID              *int         `gorm:"primaryKey"`
	UUID            string       `gorm:"unique"`
	Name            *string      `gorm:"type:varchar(255)"`
	AvatarURL       *string      `gorm:"column:avatar_url"`
	Email           *string      `gorm:"type:varchar(255)"`
	EmailVerifiedAt sql.NullTime `gorm:"default:current_timestamp"`
	Password        *string      `gorm:"type:varchar(255)"`
	SteamTradeURL   *string      `gorm:"column:steam_trade_url"`
	Experience      *int         `gorm:"column:experience"`
	Active          bool         `gorm:"column:active"`
	IsBot           bool         `gorm:"column:is_bot"`
	RememberToken   *string      `gorm:"column:remember_token"`
	CreatedAt       time.Time    `gorm:"default:current_timestamp"`
	UpdatedAt       time.Time    `gorm:"default:current_timestamp"`
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
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    *int               `bson:"user_id"`
	SteamID   *string            `bson:"steam_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
