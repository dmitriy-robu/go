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
