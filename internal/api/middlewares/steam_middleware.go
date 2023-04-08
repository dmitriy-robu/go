package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-rust-drop/internal/api/models"
	"log"
	"os"
	"time"
)

type SteamMiddleware struct {
}

var SecretKey = os.Getenv("JWT_SECRET")

func (s SteamMiddleware) VerifyToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.AbortWithStatus(401)
		return
	}

	claims, err := s.ValidateToken(token)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	c.Set("steamUserID", claims.SteamUserId)

	c.Next()
}

func (s SteamMiddleware) ValidateToken(signedToken string) (claims *models.SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &models.SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	},
	)
	if err != nil {
		log.Printf("Error while parsing token: %v\n", err)
		return
	}

	claims, ok := token.Claims.(*models.SignedDetails)
	if !ok {
		log.Printf("Invalid token: %v\n", err)
		return
	}

	timeNow := time.Now().Local().Format(time.RFC3339)

	expiredAt := claims.ExpiresAt

	if expiredAt.String() < timeNow {
		log.Printf("Token expired: %v\n", err)
		return
	}

	return claims, nil
}
