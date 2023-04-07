package models

import "github.com/golang-jwt/jwt/v5"

type SignedDetails struct {
	UserId string
	jwt.RegisteredClaims
}
