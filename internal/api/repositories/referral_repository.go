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

func (rr ReferralRepository) GetReferralTiers() ([]models.ReferralTier, error) {
	var err error
	var referralTiers []models.ReferralTier

	if err = MysqlDB.Find(&referralTiers).Error; err != nil {
		return nil, err
	}

	return referralTiers, nil
}

func (rr ReferralRepository) GetReferralByUserId(userID uint) (models.Referral, error) {
	var err error
	var referral models.Referral

	if err = MysqlDB.Where("user_id = ?", userID).First(&referral).Error; err != nil {
		return models.Referral{}, err
	}

	return referral, nil
}

func (rr ReferralRepository) GetReferralTransactionsByReferralId(referralID uint64) ([]models.ReferralTransaction, error) {
	var err error
	var referralTransactions []models.ReferralTransaction

	if err = MysqlDB.Where("referral_id = ?", referralID).Find(&referralTransactions).Error; err != nil {
		return nil, err
	}

	return referralTransactions, nil
}

func (rr ReferralRepository) GetReferralTransactionSumByReferralId(referralID uint) (int, error) {
	var err error
	var sum int

	if err = MysqlDB.Model(&models.ReferralTransaction{}).Where("referral_id = ?", referralID).Select("SUM(amount) as sum").Scan(&sum).Error; err != nil {
		return 0, err
	}

	return sum, nil
}

func (rr ReferralRepository) GetReferredUsersByUserId(userID uint) ([]models.User, error) {
	var err error
	var referredUsers []models.User

	err = MysqlDB.Model(&models.User{ID: &userID}).Association("Referrals").Find(&referredUsers)
	if err != nil {
		return nil, err
	}

	return referredUsers, nil
}
