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
	assetMaps := make([]map[string]interface{}, len(*r.referredUser))

	for i, user := range *r.referredUser {
		assetMaps[i] = map[string]interface{}{
			"name":               user.Name,
			"total_earned":       user.TotalEarnings,
			"earning_commission": user.EarningCommission,
			"current_tier":       user.CurrentTier,
			"created_at":         user.CreatedAt,
		}
	}

	return assetMaps, nil
}
