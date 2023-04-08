package repositories

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
	"time"
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

func (ur UserRepository) UpdateUserBalance(userID uint64, newBalance float64) error {
	var err error

	if err = MysqlDB.Model(&models.UserBalance{}).Where("id = ?", userID).Update("balance", newBalance).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) CreateUser(user models.User) error {
	var err error

	if err = MysqlDB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) FindUserAuthBySteamID(steamID string) (models.UserAuthSteam, error) {
	var err error
	var user models.UserAuthSteam

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return models.UserAuthSteam{}, errors.Wrap(err, "Error getting MongoDB collection")
	}

	err = collection.FindOne(ctx, bson.M{"steam_id": steamID}).Decode(&user)
	if err != nil {
		return models.UserAuthSteam{}, errors.Wrap(err, "Error finding user by steam ID")
	}

	return user, nil
}

func (ur UserRepository) UpdateUser(userID uint64) (models.User, error) {
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

/*func (ur UserRepository) GetUserIdBySteamId(steamID string) (uint64, error) {
	var err error

	return userID, nil
}*/
