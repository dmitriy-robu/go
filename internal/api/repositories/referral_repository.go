package repositories

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/requests"
	"gorm.io/gorm"
)

type ReferralRepository struct {
	MysqlDB *gorm.DB
}

func NewReferralRepository(mysql *gorm.DB) ReferralRepository {
	return ReferralRepository{
		MysqlDB: mysql,
	}
}

func (rr ReferralRepository) StoreReferralCodeToUser(user models.User, store requests.StoreUserReferralCode) (models.User, error) {
	var (
		err error
	)

	if err = rr.MysqlDB.Model(user).Update("referral_code", store.ReferralCode).Error; err != nil {
		return models.User{}, errors.Wrap(err, "Error updating user referral code")
	}

	return user, nil
}

func (rr ReferralRepository) GetReferralTiers() ([]models.ReferralTier, error) {
	var (
		err           error
		referralTiers []models.ReferralTier
	)

	if err = rr.MysqlDB.Find(&referralTiers).Error; err != nil {
		return nil, errors.Wrap(err, "Error getting referral tiers from database")
	}

	return referralTiers, nil
}

func (rr ReferralRepository) GetReferralByUserId(userID uint) (models.Referral, error) {
	var (
		err      error
		referral models.Referral
	)

	if err = rr.MysqlDB.Where("referral_user_id = ?", userID).First(&referral).Error; err != nil {
		return models.Referral{}, errors.Wrap(err, "Error getting referral by user id")
	}

	return referral, nil
}

func (rr ReferralRepository) GetReferralTransactionsByReferralId(referralID uint64) ([]models.ReferralTransaction, error) {
	var (
		err                  error
		referralTransactions []models.ReferralTransaction
	)

	if err = rr.MysqlDB.Where("referral_id = ?", referralID).Find(&referralTransactions).Error; err != nil {
		return nil, errors.Wrap(err, "Error getting referral transactions by referral id")
	}

	return referralTransactions, nil
}

func (rr ReferralRepository) GetReferralTransactionSumByReferralId(referralID uint) (uint, error) {
	var (
		err error
		sum uint
	)

	if err = rr.MysqlDB.Model(&models.ReferralTransaction{}).Where("referral_id = ?", referralID).Select("SUM(amount)").Scan(&sum).Error; err != nil {
		return 0, errors.Wrap(err, "Error getting referral transaction sum by referral id")
	}

	return sum, nil
}

func (rr ReferralRepository) GetReferredUserByUserId(userID uint) ([]models.User, error) {
	var (
		err           error
		referredUsers []models.User
	)

	if err = rr.MysqlDB.Model(&models.User{ID: userID}).Association("ReferralUsers").Find(&referredUsers); err != nil {
		return nil, errors.Wrap(err, "Error getting referred users by user id")
	}

	return referredUsers, nil
}

func (rr ReferralRepository) GetReferralTierCommissionByReferralTierLevel(level uint8) (float64, error) {
	var (
		err          error
		referralTier models.ReferralTier
	)

	if err = rr.MysqlDB.Where("level = ?", level).First(&referralTier).Error; err != nil {
		return 0.0, errors.Wrap(err, "Error getting referral tier commission by referral tier level")
	}

	return referralTier.BonusPercentage, nil
}
