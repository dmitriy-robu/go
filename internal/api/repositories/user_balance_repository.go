package repositories

import (
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type UserBalanceRepository struct {
}

func (ubr UserBalanceRepository) GetUserBalanceByUserId(userID uint64) (models.UserBalance, error) {
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

func (ubr UserBalanceRepository) CreateUserBalance(userID uint) error {
	var (
		err         error
		userBalance models.UserBalance
	)

	userBalance = models.UserBalance{
		UserID:  userID,
		Balance: 0,
	}

	if err = MysqlDB.Create(&userBalance).Error; err != nil {
		return err
	}

	return nil
}
