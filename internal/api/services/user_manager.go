package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/requests"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserManager struct {
	userRepository        repositories.UserRepository
	userBalanceRepository repositories.UserBalanceRepository
	levelRepository       repositories.LevelRepository
}

func (us UserManager) CreateOrUpdateSteamUser(userGoth goth.User) (string, error) {
	var (
		err           error
		user          models.User
		now           time.Time
		userAuth      models.UserAuthSteam
		newUUID       uuid.UUID
		userAuthSteam models.UserAuthSteam
	)

	us.userRepository.MysqlDB = MysqlDB
	us.userBalanceRepository.MysqlDB = MysqlDB

	now = time.Now()

	user = models.User{
		AvatarURL: &userGoth.AvatarURL,
		Name:      &userGoth.NickName,
		UpdatedAt: now,
	}

	userAuthSteam = models.UserAuthSteam{
		SteamUserID: userGoth.UserID,
		UpdatedAt:   now,
	}

	userAuth, err = us.userRepository.FindUserAuthBySteamID(userGoth.UserID)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return "", errors.Wrap(err, "Error finding user by SteamID")
		}

		newUUID, err = uuid.NewRandom()
		if err != nil {
			return "", errors.Wrap(err, "Error generating UUID")
		}

		user.UUID = newUUID.String()
		user.CreatedAt = now

		if user, err = us.userRepository.CreateUser(user); err != nil {
			return "", errors.Wrap(err, "Error creating user")
		}

		if err = us.userBalanceRepository.CreateUserBalance(user.ID); err != nil {
			return "", errors.Wrap(err, "Error creating user balance")
		}

		userAuthSteam.ID = primitive.NewObjectID()
		userAuthSteam.UserUUID = user.UUID
		userAuthSteam.CreatedAt = now

		if err = us.userRepository.CreateUserAuth(userAuthSteam); err != nil {
			return "", errors.Wrap(err, "Error creating user auth")
		}
	} else {
		if user, err = us.userRepository.GetUserByUuid(userAuth.UserUUID); err != nil {
			return "", errors.Wrap(err, "Error getting user from database")
		}

		if user, err = us.userRepository.UpdateUser(user); err != nil {
			return "", errors.Wrap(err, "Error updating user")
		}

		userAuthSteam.UserUUID = user.UUID

		if err = us.userRepository.UpdateUserAuth(userAuthSteam); err != nil {
			return "", errors.Wrap(err, "Error updating user auth")
		}
	}

	return user.UUID, nil
}

func (us UserManager) GetUserById(userID uint) (models.User, error) {
	var (
		err     error
		getUser models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	getUser, err = us.userRepository.FindUserByID(userID)
	if err != nil {
		return getUser, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return getUser, nil
}

func (us UserManager) AuthUser(c *gin.Context) (models.User, error) {
	var (
		err  error
		user models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userUuid, ok := c.MustGet("userUuid").(string)
	if !ok {
		return user, fmt.Errorf("user not found in context")
	}

	user, err = us.userRepository.GetUserByUuid(userUuid)
	if err != nil {
		return user, errors.Wrap(err, "Error getting user from database")
	}

	return user, nil
}

func (us UserManager) GetUserWithBalance(user models.User) (models.User, error) {
	var (
		err             error
		userWithBalance models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userWithBalance, err = us.userRepository.GetUserByIdWithBalance(user.ID)
	if err != nil {
		return userWithBalance, errors.Wrap(err, "An error occurred while retrieving user information")
	}

	return userWithBalance, nil
}

func (us UserManager) StoreSteamTradeURL(user models.User, store requests.StoreUserSteamTradeURL) error {
	var err error

	us.userRepository.MysqlDB = MysqlDB

	if user.ReferralCode != nil {
		return errors.New("Referral code already exists")
	}

	if err = us.userRepository.StoreSteamTradeURLToUser(user, store.URL); err != nil {
		return errors.Wrap(err, "Error storing referral code to user")
	}

	return nil
}
