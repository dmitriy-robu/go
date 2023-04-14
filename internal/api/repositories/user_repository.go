package repositories

import (
	"context"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mongodb"
	"go-rust-drop/internal/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
}

func (ur UserRepository) FindUserByIDWithBalance(userid int) (models.UserWithBalance, error) {
	var err error
	var userWithBalance models.UserWithBalance

	err = MysqlDB.Preload("UserWithBalance").First(&userWithBalance, userid).Error
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "Error finding user with balance")
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

func (ur UserRepository) UpdateUserBalance(userID uint64, newBalance float64) error {
	var err error

	if err = MysqlDB.Model(&models.UserBalance{}).Where("id = ?", userID).Update("balance", newBalance).Error; err != nil {
		return err
	}

	return nil
}

func (ur UserRepository) CreateUser(user models.User) (models.User, error) {
	var err error

	if err = MysqlDB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur UserRepository) UpdateUser(user models.User) (models.User, error) {
	var err error

	if err = MysqlDB.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserRepository) FindUserAuthBySteamID(steamID string) (models.UserAuthSteam, error) {
	var err error
	var userAuth models.UserAuthSteam

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return userAuth, errors.Wrap(err, "Error getting MongoDB collection")
	}

	if err = collection.FindOne(ctx, bson.M{"steam_user_id": steamID}).Decode(&userAuth); err != nil {
		if err == mongo.ErrNoDocuments {
			return userAuth, mongo.ErrNoDocuments
		}
		return userAuth, errors.Wrap(err, "Error finding user by steamID")
	}

	return userAuth, nil
}

func (ur UserRepository) CreateUserAuth(userAuthSteam models.UserAuthSteam) error {
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

func (ur UserRepository) UpdateUserAuth(userAuthSteam models.UserAuthSteam) error {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err := mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.ReplaceOne(ctx, bson.M{"steam_id": userAuthSteam.SteamUserID}, userAuthSteam)
	if err != nil {
		return errors.Wrap(err, "Error updating UserAuthSteam in MongoDB")
	}

	return nil
}

func (ur UserRepository) GetUserByUuid(uuid string) (models.User, error) {
	var err error
	var user models.User

	if err = MysqlDB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		}
		return user, err
	}

	return user, nil
}
