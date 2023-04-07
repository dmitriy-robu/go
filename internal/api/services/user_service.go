package services

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/dababase/mongodb"
	"go-rust-drop/internal/api/dababase/mysql"
	"go-rust-drop/internal/api/models"
	"strconv"
	"time"
)

type UserService struct{}

func (us UserService) CreateSteamUser(userInfo models.UserSteamInfo) (string, error) {
	db, err := mysql.GetMySQLConnection()
	if err != nil {
		return "", errors.Wrap(err, "Error getting MySQL connection")
	}

	query := "INSERT INTO users (uuid, avatar_url, name, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	//generate uuid
	uui := "123"
	createdAt := time.Now()
	updatedAt := time.Now()

	result, err := db.Exec(query, uui, userInfo.AvatarURL, userInfo.Name, createdAt, updatedAt)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting user into MySQL")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return "", errors.Wrap(err, "Error getting last insert ID from MySQL")
	}

	return strconv.FormatInt(userID, 10), nil
}

func (us UserService) InsertUserAuthSteam(userID string, steamID string) error {
	userAuthSteam := models.UserAuthSteam{
		UserID:  userID,
		SteamID: steamID,
	}

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.InsertOne(context.Background(), userAuthSteam)
	if err != nil {
		return errors.Wrap(err, "Error inserting UserAuthSteam into MongoDB")
	}

	return nil
}
