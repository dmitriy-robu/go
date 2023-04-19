package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type ReferralUserResource struct {
	referredUser *[]models.ReferredUser
	moneyConvert utils.MoneyConvert
}

func (r *ReferralUserResource) ToJSON() ([]map[string]interface{}, error) {
	var (
		referralUser  map[string]interface{}
		referralUsers []map[string]interface{}
	)

	for _, user := range *r.referredUser {
		referralUser = map[string]interface{}{
			"name":               user.Name,
			"total_earned":       user.TotalEarnings,
			"earning_commission": user.EarningCommission,
			"current_tier":       user.CurrentTier,
			"created_at":         user.CreatedAt,
		}

		referralUsers = append(referralUsers, referralUser)
	}

	return referralUsers, nil
}
