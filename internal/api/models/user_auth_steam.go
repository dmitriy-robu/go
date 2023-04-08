package models

import "github.com/golang-jwt/jwt/v5"

type SignedDetails struct {
	SteamUserId string
	Uuid        string
	jwt.RegisteredClaims
}
