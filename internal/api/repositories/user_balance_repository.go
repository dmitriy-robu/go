package repositories

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type UserBalanceRepository struct {
	MysqlDB *gorm.DB
}

func NewUserBalanceRepository(mysql *gorm.DB) UserBalanceRepository {
	return UserBalanceRepository{
		MysqlDB: mysql,
	}
}

func (ubr UserBalanceRepository) GetUserBalanceByUserId(userID uint64) (models.UserBalance, error) {
	var err error
	var userBalance models.UserBalance

	if err = ubr.MysqlDB.Where("id = ?", userID).First(&userBalance).Error; err != nil {
		return userBalance, errors.Wrap(err, "Error finding user balance")
	}

	return userBalance, nil
}

func (ubr UserBalanceRepository) CreateUserBalance(userID uint) error {
	var (
		err         error
		userBalance models.UserBalance
	)

	userBalance = models.UserBalance{
		UserID:  userID,
		Balance: 0,
	}

	if err = ubr.MysqlDB.Create(&userBalance).Error; err != nil {
		return errors.Wrap(err, "Error creating user balance")
	}

	return nil
}

func (ubr UserBalanceRepository) UpdateUserBalance(userBalance models.UserBalance) error {
	var err error

	if err = ubr.MysqlDB.Save(&userBalance).Error; err != nil {
		return errors.Wrap(err, "Error updating user balance")
	}

	return nil
}
