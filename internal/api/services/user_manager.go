package services

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/utils"
	"net/http"
	"time"
)

type UserManager struct {
	userRepository        repositories.UserRepository
	userBalanceRepository repositories.UserBalanceRepository
	levelRepository       repositories.LevelRepository
}

func (us UserManager) CreateOrUpdateSteamUser(userGoth goth.User) (string, *utils.Errors) {
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
		AvatarURL: userGoth.AvatarURL,
		Name:      userGoth.NickName,
		UpdatedAt: now,
	}

	userAuthSteam = models.UserAuthSteam{
		SteamUserID: userGoth.UserID,
		UpdatedAt:   now,
	}

	userAuth, err = us.userRepository.FindUserAuthBySteamID(userGoth.UserID)

	if err != nil {
		newUUID, _ = uuid.NewRandom()

		user.UUID = newUUID.String()
		user.CreatedAt = now

		if user, err = us.userRepository.CreateUser(user); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error creating user", err)
		}

		if err = us.userBalanceRepository.CreateUserBalance(user.ID); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error creating user balance", err)
		}

		userAuthSteam.UserUUID = user.UUID
		userAuthSteam.CreatedAt = now

		if err = us.userRepository.CreateUserAuth(userAuthSteam); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error creating user", err)
		}
	} else {
		if user, err = us.userRepository.GetUserByUuid(userAuth.UserUUID); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error getting user", err)
		}

		if user, err = us.userRepository.UpdateUser(user); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error updating user", err)
		}

		userAuthSteam.UserUUID = user.UUID

		if err = us.userRepository.UpdateUserAuth(userAuthSteam); err != nil {
			return "", utils.NewErrors(http.StatusInternalServerError, "Error updating user auth", err)
		}
	}

	return user.UUID, nil
}

func (us UserManager) GetUserById(userID uint) (models.User, *utils.Errors) {
	var (
		err  error
		user models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	user, err = us.userRepository.FindUserByID(userID)
	if err != nil {
		return user, utils.NewErrors(http.StatusInternalServerError, "An error occurred while retrieving user information", err)
	}

	return user, nil
}

func (us UserManager) AuthUser(c *gin.Context) (models.User, *utils.Errors) {
	var (
		err  error
		user models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userUuid, ok := c.MustGet("userUuid").(string)
	if !ok {
		return user, utils.NewErrors(http.StatusUnauthorized, "Unauthorized", errors.New("Unauthorized"))
	}

	user, err = us.userRepository.GetUserByUuid(userUuid)
	if err != nil {
		return user, utils.NewErrors(http.StatusUnauthorized, "Unauthorized", err)
	}

	return user, nil
}

func (us UserManager) GetUserWithBalance(user models.User) (models.User, *utils.Errors) {
	var (
		err             error
		userWithBalance models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userWithBalance, err = us.userRepository.GetUserByIdWithBalance(user.ID)
	if err != nil {
		return userWithBalance, utils.NewErrors(http.StatusInternalServerError, "An error occurred while retrieving user information", err)
	}

	return userWithBalance, nil
}

func (us UserManager) StoreSteamTradeURL(user models.User, store requests.StoreUserSteamTradeURL) *utils.Errors {
	var (
		err error
	)

	us.userRepository.MysqlDB = MysqlDB

	if user.ReferralCode != "" {
		return utils.NewErrors(http.StatusBadRequest, "You already have a referral code", err)
	}

	if err = us.userRepository.StoreSteamTradeURLToUser(user, store.URL); err != nil {
		return utils.NewErrors(http.StatusInternalServerError, "An error occurred while storing Steam trade URL", err)
	}

	return nil
}
