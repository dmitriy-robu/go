package repositories

import (
	"go-rust-drop/internal/api/mappers"
	"go-rust-drop/internal/api/models"
)

type ReferralRepository struct {
}

func (rr ReferralRepository) StoreReferralCodeToUser(user *models.User, store *mappers.StoreUserReferralCode) (*models.User, error) {
	var err error

	if err = MysqlDB.Model(user).Update("referral_code", store.ReferralCode).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}
