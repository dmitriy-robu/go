package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
)

type ReferralUserResource struct {
	User               *[]models.User
	moneyConvert       utils.MoneyConvert
	error              utils.Errors
	referralRepository repositories.ReferralRepository
}

func (r *ReferralUserResource) ToJSON() ([]map[string]interface{}, error) {
	assetMaps := make([]map[string]interface{}, len(*r.User))

	for i, user := range *r.User {
		assetMaps[i] = map[string]interface{}{
			"name":               user.Name,
			"total_earned":       r.moneyConvert.FromCentsToVault(r.getSum(user.ID)), // Здесь нужно заменить на сумму реферальных транзакций пользователя
			"earning_commission": r.getCommission(&user.ReferralTierLevel),
			"current_tier":       user.ReferralTierLevel,
			"created_at":         user.CreatedAt,
		}
	}

	return assetMaps, nil
}

func (r *ReferralUserResource) getCommission(level *uint) float64 {
	commission, err := r.referralRepository.GetReferralTierCommissionByReferralTierLevel(*level)
	if err != nil {
		r.error.ResourcesHandleError("Referral tier commission not found", err)
		return 0
	}

	return commission
}

func (r *ReferralUserResource) getSum(userID *uint) int {
	referral, err := r.referralRepository.GetReferralByUserId(*userID)
	if err != nil {
		r.error.ResourcesHandleError("Referral not found", err)
		return 0
	}

	sum, err := r.referralRepository.GetReferralTransactionSumByReferralId(referral.ID)
	if err != nil {
		r.error.ResourcesHandleError("Referral transaction sum not found", err)
		return 0
	}

	return sum
}
