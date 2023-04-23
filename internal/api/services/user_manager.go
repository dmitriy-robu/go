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

func (us UserManager) CreateOrUpdateSteamUser(userGoth goth.User) (string, utils.Errors) {
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
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error creating user",
				Err:     err,
			}
		}

		if err = us.userBalanceRepository.CreateUserBalance(user.ID); err != nil {
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error creating user balance",
				Err:     err,
			}
		}

		userAuthSteam.UserUUID = user.UUID
		userAuthSteam.CreatedAt = now

		if err = us.userRepository.CreateUserAuth(userAuthSteam); err != nil {
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error creating user",
				Err:     err,
			}
		}
	} else {
		if user, err = us.userRepository.GetUserByUuid(userAuth.UserUUID); err != nil {
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error getting user",
				Err:     err,
			}
		}

		if user, err = us.userRepository.UpdateUser(user); err != nil {
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error updating user",
				Err:     err,
			}
		}

		if err = us.userRepository.UpdateUserAuth(userAuthSteam); err != nil {
			return "", utils.Errors{
				Code:    http.StatusInternalServerError,
				Message: "Error updating user auth",
				Err:     err,
			}
		}
	}

	return user.UUID, utils.Errors{}
}

func (us UserManager) GetUserById(userID uint) (models.User, utils.Errors) {
	var (
		err  error
		user models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	user, err = us.userRepository.FindUserByID(userID)
	if err != nil {
		return user, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "An error occurred while retrieving user information",
			Err:     err,
		}
	}

	return user, utils.Errors{}
}

func (us UserManager) AuthUser(c *gin.Context) (models.User, utils.Errors) {
	var (
		err  error
		user models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userUuid, ok := c.MustGet("userUuid").(string)
	if !ok {
		return user, utils.Errors{
			Code:    401,
			Message: "Unauthorized",
			Err:     errors.New("Unauthorized"),
		}
	}

	user, err = us.userRepository.GetUserByUuid(userUuid)
	if err != nil {
		return user, utils.Errors{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
			Err:     err,
		}
	}

	return user, utils.Errors{}
}

func (us UserManager) GetUserWithBalance(user models.User) (models.User, utils.Errors) {
	var (
		err             error
		userWithBalance models.User
	)

	us.userRepository.MysqlDB = MysqlDB

	userWithBalance, err = us.userRepository.GetUserByIdWithBalance(user.ID)
	if err != nil {
		return userWithBalance, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "An error occurred while retrieving user information",
			Err:     err,
		}
	}

	return userWithBalance, utils.Errors{}
}

func (us UserManager) StoreSteamTradeURL(user models.User, store requests.StoreUserSteamTradeURL) utils.Errors {
	var (
		err error
	)

	us.userRepository.MysqlDB = MysqlDB

	if user.ReferralCode != "" {
		return utils.Errors{
			Code:    http.StatusBadRequest,
			Message: "You already have a referral code",
			Err:     err,
		}
	}

	if err = us.userRepository.StoreSteamTradeURLToUser(user, store.URL); err != nil {
		return utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "An error occurred while storing Steam trade URL",
			Err:     err,
		}
	}

	return utils.Errors{}
}
