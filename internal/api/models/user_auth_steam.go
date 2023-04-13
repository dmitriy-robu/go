package models

import "github.com/golang-jwt/jwt/v5"

type SignedDetails struct {
	SteamUserId string
	Uuid        string
	jwt.RegisteredClaims
}

type Session struct {
	ID        string
	UserID    string
	Email     string
	Name      string
	AvatarURL string
	Provider  string
}
