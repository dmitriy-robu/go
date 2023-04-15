package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func (us UserService) CreateOrUpdateSteamUser(userGoth goth.User) (string, error) {
	now := time.Now()

	user := models.User{
		AvatarURL: &userGoth.AvatarURL,
		Name:      &userGoth.NickName,
		UpdatedAt: now,
	}

	userAuth, err := us.userRepo.FindUserAuthBySteamID(userGoth.UserID)

	if err != nil && err != mongo.ErrNoDocuments {
		return "", errors.Wrap(err, "Error finding user by SteamID")
	}

	userAuthSteam := models.UserAuthSteam{
		SteamUserID: &userGoth.UserID,
		UpdatedAt:   now,
	}

	if err == nil {
		user, err = us.userRepo.GetUserByUuid(userAuth.UserUUID)
		if err != nil {
			return "", errors.Wrap(err, "Error getting user from database")
		}

		user, err = us.userRepo.UpdateUser(user)
		if err != nil {
			return "", errors.Wrap(err, "Error updating user")
		}

		userAuthSteam.UserUUID = user.UUID

		if err = us.userRepo.UpdateUserAuth(userAuthSteam); err != nil {
			return "", errors.Wrap(err, "Error updating user auth")
		}
	} else {
		newUUID, err := uuid.NewRandom()
		if err != nil {
			return "", errors.Wrap(err, "Error generating UUID")
		}

		user.UUID = newUUID.String()
		user.CreatedAt = now

		user, err = us.userRepo.CreateUser(user)
		if err != nil {
			return "", errors.Wrap(err, "Error creating user")
		}

		userAuthSteam.UserUUID = user.UUID
		userAuthSteam.CreatedAt = now

		if err = us.userRepo.CreateUserAuth(userAuthSteam); err != nil {
			return "", errors.Wrap(err, "Error creating user auth")
		}
	}

	return user.UUID, nil
}

func (us UserService) GetUser(user models.User) (models.User, error) {
	var err error

	userWithBalance, err := us.userRepo.FindUserByID(*user.ID)
	if err != nil {
		return models.User{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}

func (us UserService) AuthUser(c *gin.Context) (user models.User, err error) {
	userUuid, ok := c.MustGet("userUuid").(string)
	if !ok {
		return models.User{}, fmt.Errorf("user not found in context")
	}

	user, err = us.userRepo.GetUserByUuid(userUuid)
	if err != nil {
		return models.User{}, errors.Wrap(err, "Error getting user from database")
	}

	return user, nil
}

func (us UserService) GetUserWithBalance(user models.User) (models.User, error) {
	var err error

	userWithBalance, err := us.userRepo.GetUserByIdWithBalance(*user.ID)
	if err != nil {
		return models.User{}, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}
