package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	userRepo  repositories.UserRepository
	db        *gorm.DB
	userModel models.User
}

func (us UserService) CreateSteamUser(userInfo models.UserSteamInfo) (models.User, error) {
	var err error

	if *userInfo.SteamID == "" {
		return us.userModel, errors.New("SteamID is empty")
	}

	userAuth, err := us.userRepo.FindUserAuthBySteamID(*userInfo.SteamID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return us.userModel, errors.Wrap(err, "Error finding user by SteamID")
		}
	}

	if userAuth.UserID != nil {
		user, err := us.userRepo.UpdateUser(*userAuth.UserID)
		if err != nil {
			return us.userModel, errors.Wrap(err, "Error finding user by ID")
		}

		return user, nil
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return us.userModel, errors.Wrap(err, "Error generating UUID")
	}

	user := models.User{
		UUID:      newUUID.String(),
		AvatarURL: userInfo.AvatarURL,
		Name:      userInfo.Name,
	}

	if err = us.userRepo.CreateUser(user); err != nil {
		return us.userModel, errors.Wrap(err, "Error creating user")
	}

	return user, nil
}

func (us UserService) CreateUserAuthSteam(userAuthSteam models.UserAuthSteam) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.InsertOne(ctx, userAuthSteam)
	if err != nil {
		return errors.Wrap(err, "Error inserting UserAuthSteam into MongoDB")
	}

	return nil
}

func (us UserService) GetUserInfo(steamUserID string) (models.UserWithBalance, error) {
	var err error

	userWithBalance, err := us.userRepo.FindUserByIDWithBalance(steamUserID)
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}
