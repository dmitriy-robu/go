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

func (ur UserRepository) FindUserByID(userID uint) (models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = MysqlDB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.Wrap(err, "User not found")
		}
		return user, errors.Wrap(err, "Error finding user by ID")
	}

	return user, nil
}

func (ur UserRepository) UpdateUserBalance(userID uint64, newBalance float64) error {
	var err error

	if err = MysqlDB.Model(&models.UserBalance{}).Where("user_id = ?", userID).Update("balance", newBalance).Error; err != nil {
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
		return user, errors.Wrap(err, "Error updating user")
	}

	return user, nil
}

func (ur UserRepository) FindUserAuthBySteamID(steamID string) (models.UserAuthSteam, error) {
	var (
		err        error
		userAuth   models.UserAuthSteam
		collection *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("user_auth_steam")
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

func (ur UserRepository) GetUserAuthByUserUUID(uuid string) (models.UserAuthSteam, error) {
	var (
		err        error
		userAuth   models.UserAuthSteam
		collection *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return userAuth, errors.Wrap(err, "Error getting MongoDB collection")
	}

	if err = collection.FindOne(ctx, bson.M{"user_uuid": uuid}).Decode(&userAuth); err != nil {
		if err == mongo.ErrNoDocuments {
			return userAuth, mongo.ErrNoDocuments
		}
		return userAuth, errors.Wrap(err, "Error finding user by steamID")
	}

	return userAuth, nil
}

func (ur UserRepository) CreateUserAuth(userAuthSteam models.UserAuthSteam) error {
	var (
		err        error
		collection *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("user_auth_steam")
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
	var (
		err        error
		collection *mongo.Collection
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection, err = mongodb.GetCollectionByName("user_auth_steam")
	if err != nil {
		return errors.Wrap(err, "Error getting MongoDB collection")
	}

	_, err = collection.ReplaceOne(ctx, bson.M{"steam_user_id": userAuthSteam.SteamUserID}, userAuthSteam)
	if err != nil {
		return errors.Wrap(err, "Error updating UserAuthSteam in MongoDB")
	}

	return nil
}

func (ur UserRepository) GetUserByUuid(uuid string) (models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = MysqlDB.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.Wrap(err, "User not found")
		}
		return user, errors.Wrap(err, "Error finding user by uuid")
	}

	return user, nil
}

func (ur UserRepository) GetUserByIdWithBalance(userID uint) (models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = MysqlDB.Preload("UserBalance").Where("id = ?", userID).First(&user).Error; err != nil {
		return user, errors.Wrap(err, "Error finding user with balance")
	}

	return user, nil
}

func (ur UserRepository) StoreSteamTradeURLToUser(user models.User, steamTradeURL string) error {
	var err error

	if err = MysqlDB.Model(user).Update("steam_trade_url", steamTradeURL).Error; err != nil {
		return errors.Wrap(err, "Error updating user with steam trade url")
	}

	return nil
}
