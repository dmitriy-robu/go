package services

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/dababase/mongodb"
	"go-rust-drop/internal/api/dababase/mysql"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repository"
	"strconv"
	"time"
)

type UserService struct{}

func (us UserService) CreateSteamUser(userInfo models.UserSteamInfo) (string, error) {
	db, err := mysql.GetMySQLConnection()
	if err != nil {
		return "", errors.Wrap(err, "Error getting MySQL connection")
	}

	ds := goqu.Insert("users").Rows(
		goqu.Record{
			"uuid":       "123",
			"avatar_url": userInfo.AvatarURL,
			"name":       userInfo.Name,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		},
	)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return "", errors.Wrap(err, "Error building SQL query")
	}

	result, err := db.Exec(sql)
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

func (us UserService) GetUserInfo(userID uint64) (models.UserWithBalance, error) {
	db, err := mysql.GetMySQLConnection()

	userWithBalance, err := repository.UserRepository{}.FindUserByIDWithBalance(db, userID)
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}
