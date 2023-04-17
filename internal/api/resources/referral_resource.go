package resources

import (
	"encoding/json"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type ReferralDetailResource struct {
	ReferralDetails map[string]interface{}
	moneyConvert    utils.MoneyConvert
}

func (r *ReferralDetailResource) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"referral_code":  r.ReferralDetails["referral_code"],
		"total_earnings": r.moneyConvert.FromCentsToVault(r.ReferralDetails["total_earnings"].(int)),
		"referred_users": r.ReferralDetails["referred_users"], // Здесь нужно преобразовать каждого пользователя в ReferralUserResource
	})
}

type ReferralUserResource struct {
	Referral     *models.User
	moneyConvert utils.MoneyConvert
}

func (r *ReferralUserResource) ToJSON() ([]byte, error) {
	//authorizedUser := AuthUser() // Здесь нужно получить авторизованного пользователя, возможно, из контекста

	return json.Marshal(map[string]interface{}{
		"name": r.Referral.Name,
		//"total_earned":       r.moneyConvert.FromCentsToVault(r.Referral.ReferralTransactionsSum), // Здесь нужно заменить на сумму реферальных транзакций пользователя
		//"earning_commission": authorizedUser.ReferralTierLevel,                                    // Здесь нужно заменить на соответствующее значение
		"current_tier": r.Referral.ReferralTierLevel,
		"created_at":   r.Referral.CreatedAt,
	})
}
