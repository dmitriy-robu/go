package repositories

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/request"
)

type ReferralRepository struct {
}

func (rr ReferralRepository) StoreReferralCodeToUser(user *models.User, store *request.StoreUserReferralCode) (*models.User, error) {
	var err error

	if err = MysqlDB.Model(user).Update("referral_code", store.ReferralCode).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}
