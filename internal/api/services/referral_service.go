package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/requests"
)

type ReferralManager struct {
	referralRepository repositories.ReferralRepository
}

func (rs ReferralManager) StoreReferralCode(user models.User, store requests.StoreUserReferralCode) (models.User, error) {
	var err error

	if user.ReferralCode != nil {
		return models.User{}, errors.New("Referral code already exists")
	}

	user, err = rs.referralRepository.StoreReferralCodeToUser(user, store)
	if err != nil {
		return models.User{}, errors.Wrap(err, "Error storing referral code to user")
	}

	return user, nil
}

func (rs ReferralManager) GetReferralDetails(user models.User) (models.ReferralDetails, error) {
	var (
		err                   error
		referralTiers         []models.ReferralTier
		referral              models.Referral
		totalEarnings         int
		referredUsers         []models.ReferredUser
		referralDetails       models.ReferralDetails
		currentTierCommission float64
	)

	referralTiers, err = rs.referralRepository.GetReferralTiers()
	if err != nil {
		return referralDetails, errors.Wrap(err, "Error getting referral tiers from repository")
	}

	currentTierCommission = 0.0
	if user.ReferralTierLevel > 0 {
		for _, tier := range referralTiers {
			if tier.Level == int(user.ReferralTierLevel) {
				currentTierCommission = tier.BonusPercentage
				break
			}
		}
	}

	referral, err = rs.referralRepository.GetReferralByUserId(user.ID)
	if err != nil {
		return referralDetails, errors.Wrap(err, "Error getting referral by user id")
	}

	totalEarnings, err = rs.referralRepository.GetReferralTransactionSumByReferralId(referral.ID)
	if err != nil {
		return referralDetails, errors.Wrap(err, "Error getting referral transaction sum")
	}

	referredUsers, err = rs.getReferredUsers(referral.ReferralUserID, referral.ID)
	if err != nil {
		return referralDetails, errors.Wrap(err, "Error getting referred users by user id")
	}

	referralDetails = models.ReferralDetails{
		ReferralCode:          user.ReferralCode,
		TotalEarnings:         totalEarnings,
		CurrentTierCommission: currentTierCommission,
		ReferredUsers:         referredUsers,
	}

	return referralDetails, nil
}

func (rs ReferralManager) getReferredUsers(userID uint, referralID uint) ([]models.ReferredUser, error) {
	var (
		err           error
		users         []models.User
		referredUser  models.ReferredUser
		referredUsers []models.ReferredUser
		commission    float64
	)

	users, err = rs.referralRepository.GetReferredUserByUserId(userID)
	if err != nil {
		return referredUsers, errors.Wrap(err, "Error getting referred users by user id")
	}

	for _, user := range users {
		commission, err = rs.referralRepository.GetReferralTierCommissionByReferralTierLevel(user.ReferralTierLevel)
		if err != nil {
			commission = 0.0
		}

		sum, err := rs.referralRepository.GetReferralTransactionSumByReferralId(referralID)
		if err != nil {
			sum = 0
		}

		referredUser = models.ReferredUser{
			Name:              *user.Name,
			TotalEarnings:     sum,
			EarningCommission: commission,
			CurrentTier:       user.ReferralTierLevel,
			CreatedAt:         user.CreatedAt,
		}

		referredUsers = append(referredUsers, referredUser)
	}

	return referredUsers, nil
}
