package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/request"
)

type ReferralService struct {
	referralRepository repositories.ReferralRepository
}

func (rs ReferralService) StoreReferralCode(user *models.User, store *request.StoreUserReferralCode) (*models.User, error) {
	var err error

	if user.ReferralCode != nil {
		return &models.User{}, errors.New("Referral code already exists")
	}

	user, err = rs.referralRepository.StoreReferralCodeToUser(user, store)
	if err != nil {
		return &models.User{}, errors.Wrap(err, "Error storing referral code to user")
	}

	return user, nil
}

func (rs ReferralService) GetReferralDetails(user models.User) (map[string]interface{}, error) {
	var (
		err           error
		referralTiers []models.ReferralTier
		referral      models.Referral
		totalEarnings int
		referredUsers []models.User
	)

	referralTiers, err = rs.referralRepository.GetReferralTiers()
	if err != nil {
		return nil, errors.Wrap(err, "Error getting referral tiers from repository")
	}

	currentTierCommission := 0.0
	if user.ReferralTierLevel > 0 {
		for _, tier := range referralTiers {
			if tier.Level == int(user.ReferralTierLevel) {
				currentTierCommission = tier.BonusPercentage
				break
			}
		}
	}

	referral, err = rs.referralRepository.GetReferralByUserId(*user.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting referral by user id")
	}

	totalEarnings, err = rs.referralRepository.GetReferralTransactionSumByReferralId(referral.ID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting referral transaction sum")
	}

	referredUsers, err = rs.referralRepository.GetReferredUsersByUserId(referral.ReferralUserID)
	if err != nil {
		return nil, errors.Wrap(err, "Error getting referred users by user id")
	}

	referralDetails := map[string]interface{}{
		"referral_code":           user.ReferralCode,
		"total_earnings":          totalEarnings,
		"current_tier_commission": currentTierCommission,
		"referred_users":          referredUsers,
	}

	return referralDetails, nil
}
