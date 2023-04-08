package repositories

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/database/mysql"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	userWithBalance models.UserWithBalance
}

func (u UserRepository) FindUserByIDWithBalance(userID uint64) (models.UserWithBalance, error) {
	var err error

	err = MysqlDB.Preload("UserBalance").First(&u.userWithBalance, userID).Error
	if err != nil {
		return models.UserWithBalance{}, errors.Wrap(err, "Error finding user with balance")
	}

	userWithBalance := models.UserWithBalance{
		User:        u.userWithBalance.User,
		UserBalance: u.userWithBalance.UserBalance,
	}

	return userWithBalance, nil
}

func (u UserRepository) FindUserByID(userID uint64) (models.User, error) {
	var err error
	var db *gorm.DB

	db, err = mysql.GetGormConnection()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return user, nil
}

func (u UserRepository) GetUserBalance(db *sql.DB, userID uint64) (models.UserBalance, error) {
	var err error

	ds := goqu.From(models.TableUserBalance).Select(
		"id",
		"user_id",
		"balance",
	).Where(
		goqu.Ex{"user_id": userID},
	)

	query, _, err := ds.ToSQL()
	if err != nil {
		return models.UserBalance{}, err
	}

	row := db.QueryRow(query, userID)

	var userBalance models.UserBalance
	err = row.Scan(&userBalance.ID, &userBalance.UserID, &userBalance.Balance)
	if err != nil {
		return models.UserBalance{}, err
	}

	return userBalance, nil
}
