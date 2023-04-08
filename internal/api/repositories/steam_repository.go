package repositories

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type SteamRepository struct {
}

func (sr SteamRepository) UpdateAllTokens(signedToken, signedRefreshToken, uid string) error {
	var updateObj primitive.D

	UpdatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateObj = append(updateObj, bson.E{Key: "$set", Value: bson.D{
		{"token", signedToken},
		{"refresh_token", signedRefreshToken},
		{"updated_at", UpdatedAt}}})

	upsert := true
	filter := bson.M{"user_id": uid}
	opts := options.UpdateOptions{Upsert: &upsert}

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.UpdateOne(context.Background(), filter, updateObj, &opts)
	if err != nil {
		return err
	}

	return nil
}

func (sr SteamRepository) GenerateAllTokens(steamUserId, uuid string) (signedToken string, signedRefreshToken string, err error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	expiredTime := time.Now().Local().Add(time.Hour * time.Duration(24))
	claims := &models.SignedDetails{
		SteamUserId: steamUserId,
		Uuid:        uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	refreshClaims := &models.SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))

	if err != nil {
		log.Printf("Error while signing the token: %v", err)
		return
	}

	return token, refreshToken, err
}
