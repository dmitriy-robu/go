package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/dababase/mongodb"
	"go-rust-drop/internal/api/dababase/mysql"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserService struct {
	userRepo repositories.UserRepository
	db       *gorm.DB
}

func (us UserService) CreateSteamUser(userInfo models.UserSteamInfo) (string, error) {
	db, err := mysql.GetGormConnection()
	if err != nil {
		return "", errors.Wrap(err, "Error getting MySQL connection")
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "Error generating UUID")
	}

	user := models.User{
		UUID:      newUUID.String(),
		AvatarURL: userInfo.AvatarURL,
		Name:      userInfo.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		return "", errors.Wrap(err, "Error inserting user into MySQL")
	}

	return strconv.FormatUint(*user.ID, 10), nil
}

func (us UserService) InsertUserAuthSteam(userID string, steamID string) error {
	var err error

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
	var err error

	userWithBalance, err := us.userRepo.FindUserByIDWithBalance(userID)
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}
