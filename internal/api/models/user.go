package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User полная таблица в mysql
type User struct {
	ID              uint64     `json:"id"`
	UUID            string     `json:"uuid"`
	Name            string     `json:"name"`
	AvatarURL       *string    `json:"avatar_url,omitempty"`
	Email           *string    `json:"email,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	Password        *string    `json:"-"`
	SteamTradeURL   *string    `json:"steam_trade_url,omitempty"`
	Experience      *uint64    `json:"experience,omitempty"`
	Active          bool       `json:"active"`
	IsBot           bool       `json:"is_bot"`
	RememberToken   *string    `json:"remember_token,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

type UserInfo struct {
	ID            int
	UUID          string
	Name          string
	AvatarURL     string
	Balance       int
	SteamTradeURL string
	Experience    int
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
