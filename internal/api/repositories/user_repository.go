package repositories

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	userWithBalance models.UserWithBalance
}

func (ur UserRepository) FindUserByIDWithBalance(steamUserID string) (models.UserWithBalance, error) {
	var err error

	err = MysqlDB.Preload("UserBalance").First(&ur.userWithBalance, steamUserID).Error
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "Error finding user with balance")
	}

	userWithBalance := models.UserWithBalance{
		User:        ur.userWithBalance.User,
		UserBalance: ur.userWithBalance.UserBalance,
	}

	return userWithBalance, nil
}

func (ur UserRepository) FindUserByID(userID uint64) (models.User, error) {
	var err error
	var user models.User

	if err = MysqlDB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return user, nil
}

func (ur UserRepository) GetUserBalance(userID uint64) (models.UserBalance, error) {
	var err error
	var userBalance models.UserBalance

	if err = MysqlDB.Where("id = ?", userID).First(&userBalance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return userBalance, err
		}
		return userBalance, err
	}

	return userBalance, nil
}

/*func (ur UserRepository) GetUserIdBySteamId(steamID string) (uint64, error) {
	var err error

	return userID, nil
}*/