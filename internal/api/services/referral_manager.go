package services

import (
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type ReferralManager struct {
	referralRepository repositories.ReferralRepository
}

func (rs ReferralManager) StoreReferralCode(user models.User, store requests.StoreUserReferralCode) (models.User, utils.Errors) {
	var (
		err error
	)

	if user.ReferralCode != "" {
		return models.User{}, utils.Errors{
			Code:    http.StatusBadRequest,
			Message: "User already has a referral code",
			Err:     errors.New("User already has a referral code"),
		}
	}

	rs.referralRepository.MysqlDB = MysqlDB

	user, err = rs.referralRepository.StoreReferralCodeToUser(user, store)
	if err != nil {
		return models.User{}, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error storing referral code to user",
			Err:     err,
		}
	}

	return user, utils.Errors{}
}

func (rs ReferralManager) GetReferralDetails(user models.User) (models.ReferralDetails, utils.Errors) {
	var (
		err                   error
		referralTiers         []models.ReferralTier
		referral              models.Referral
		totalEarnings         uint
		referredUsers         []models.ReferredUser
		referralDetails       models.ReferralDetails
		currentTierCommission float64
		errorHandler          utils.Errors
	)

	rs.referralRepository.MysqlDB = MysqlDB

	referralTiers, err = rs.referralRepository.GetReferralTiers()
	if err != nil {
		return referralDetails, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error getting referral tiers",
			Err:     err,
		}
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
		return referralDetails, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error getting referral by user id",
			Err:     err,
		}
	}

	totalEarnings, err = rs.referralRepository.GetReferralTransactionSumByReferralId(referral.ID)
	if err != nil {
		return referralDetails, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error getting referral transaction sum by referral id",
			Err:     err,
		}
	}

	referredUsers, errorHandler = rs.getReferredUsers(referral.ReferralUserID, referral.ID)
	if errorHandler.Err != nil {
		return referralDetails, errorHandler
	}

	referralDetails = models.ReferralDetails{
		ReferralCode:          user.ReferralCode,
		TotalEarnings:         totalEarnings,
		CurrentTierCommission: currentTierCommission,
		ReferredUsers:         referredUsers,
	}

	return referralDetails, errorHandler
}

func (rs ReferralManager) getReferredUsers(userID uint, referralID uint) ([]models.ReferredUser, utils.Errors) {
	var (
		err           error
		users         []models.User
		referredUser  models.ReferredUser
		referredUsers []models.ReferredUser
		commission    float64
		sum           uint
		user          models.User
	)

	rs.referralRepository.MysqlDB = MysqlDB

	users, err = rs.referralRepository.GetReferredUserByUserId(userID)
	if err != nil {
		return referredUsers, utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error getting referred users by user id",
			Err:     err,
		}
	}

	for _, user = range users {
		commission, err = rs.referralRepository.GetReferralTierCommissionByReferralTierLevel(user.ReferralTierLevel)
		if err != nil {
			commission = 0.0
		}

		sum, err = rs.referralRepository.GetReferralTransactionSumByReferralId(referralID)
		if err != nil {
			sum = 0
		}

		referredUser = models.ReferredUser{
			Name:              user.Name,
			TotalEarnings:     sum,
			EarningCommission: commission,
			CurrentTier:       user.ReferralTierLevel,
			CreatedAt:         user.CreatedAt,
		}

		referredUsers = append(referredUsers, referredUser)
	}

	return referredUsers, utils.Errors{}
}
