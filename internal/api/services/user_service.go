package services

import (
	"github.com/google/uuid"
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

func (us UserService) CreateOrUpdateSteamUser(userInfo models.UserSteamInfo) (models.User, error) {
	var err error
	var userAuthSteam models.UserAuthSteam
	var user models.User

	if *userInfo.SteamID == "" {
		return us.userModel, errors.New("SteamID is empty")
	}

	userAuth, err := us.userRepo.FindUserAuthBySteamID(*userInfo.SteamID)

	if err == nil {
		userUpdate := models.User{
			AvatarURL: userInfo.AvatarURL,
			Name:      userInfo.Name,
			UpdatedAt: time.Now(),
		}

		user, err = us.userRepo.UpdateUser(userAuth.UserID, userUpdate)
		if err != nil {
			return us.userModel, errors.Wrap(err, "Error updating user")
		}

		userAuthSteam = models.UserAuthSteam{
			SteamID:   userInfo.SteamID,
			UserID:    user.ID,
			UpdatedAt: time.Now(),
		}

		if err = us.userRepo.UpdateUserAuth(userAuthSteam); err != nil {
			return us.userModel, errors.Wrap(err, "Error updating user auth")
		}

		return user, nil
	} else if err != mongo.ErrNoDocuments {
		return us.userModel, errors.Wrap(err, "Error finding user by SteamID")
	}

	newUUID, err := uuid.NewRandom()
	if err != nil {
		return us.userModel, errors.Wrap(err, "Error generating UUID")
	}

	user = models.User{
		UUID:      newUUID.String(),
		AvatarURL: userInfo.AvatarURL,
		Name:      userInfo.Name,
	}

	user, err = us.userRepo.CreateUser(user)
	if err != nil {
		return user, errors.Wrap(err, "Error creating user")
	}

	if err = us.userRepo.CreateUserAuth(models.UserAuthSteam{
		SteamID: userInfo.SteamID,
		UserID:  user.ID,
	}); err != nil {
		return user, errors.Wrap(err, "Error creating user auth")
	}

	return user, nil
}

func (us UserService) GetUserInfo(steamUserID string) (models.UserWithBalance, error) {
	var err error

	userWithBalance, err := us.userRepo.FindUserByIDWithBalance(steamUserID)
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}
