package services

import (
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserService struct {
	userRepo  repositories.UserRepository
	userModel models.User
}

func (us UserService) CreateOrUpdateSteamUser(userGoth goth.User) error {
	now := time.Now()

	user := models.User{
		AvatarURL: &userGoth.AvatarURL,
		Name:      &userGoth.NickName,
		UpdatedAt: now,
	}

	userAuth, err := us.userRepo.FindUserAuthBySteamID(userGoth.UserID)

	if err != nil && err != mongo.ErrNoDocuments {
		return errors.Wrap(err, "Error finding user by SteamID")
	}

	userAuthSteam := models.UserAuthSteam{
		SteamUserID: &userGoth.UserID,
		UpdatedAt:   now,
	}

	if err == nil {
		user, err = us.userRepo.UpdateUser(userAuth.UserUUID, user)
		if err != nil {
			return errors.Wrap(err, "Error updating user")
		}

		userAuthSteam.UserUUID = user.UUID

		if err = us.userRepo.UpdateUserAuth(userAuthSteam); err != nil {
			return errors.Wrap(err, "Error updating user auth")
		}
	} else {
		newUUID, err := uuid.NewRandom()
		if err != nil {
			return errors.Wrap(err, "Error generating UUID")
		}

		user.UUID = newUUID.String()
		user.CreatedAt = now

		user, err = us.userRepo.CreateUser(user)
		if err != nil {
			return errors.Wrap(err, "Error creating user")
		}

		userAuthSteam.UserUUID = user.UUID
		userAuthSteam.CreatedAt = now

		if err = us.userRepo.CreateUserAuth(userAuthSteam); err != nil {
			return errors.Wrap(err, "Error creating user auth")
		}
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
