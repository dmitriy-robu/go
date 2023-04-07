package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User полная таблица в mysql
type User struct {
	ID              *uint64
	UUID            *string
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
}

type UserSteamInfo struct {
	SteamID   string
	AvatarURL string
	Name      string
}

type UserAuthSteam struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"user_id"`
	SteamID string             `bson:"steam_id"`
}

type UserSteamProfile struct {
	SteamID   string `json:"steamid"`
	AvatarURL string `json:"avatar"`
	Name      string `json:"personaname"`
}

type UserWithBalance struct {
	User        User
	UserBalance UserBalance
}
