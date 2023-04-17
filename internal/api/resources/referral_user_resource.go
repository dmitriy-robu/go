package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type ReferralUserResource struct {
	User         *[]models.User
	moneyConvert utils.MoneyConvert
}

func (r *ReferralUserResource) ToJSON() ([]map[string]interface{}, error) {
	assetMaps := make([]map[string]interface{}, len(*r.User))

	for i, user := range *r.User {
		assetMaps[i] = map[string]interface{}{
			"name": user.Name,
			//"total_earned":       r.moneyConvert.FromCentsToVault(user.ReferralTransactionsSum), // Здесь нужно заменить на сумму реферальных транзакций пользователя
			"earning_commission": user.ReferralTierLevel, // Здесь нужно заменить на соответствующее значение
			"current_tier":       user.ReferralTierLevel,
			"created_at":         user.CreatedAt,
		}
	}

	return assetMaps, nil
}
