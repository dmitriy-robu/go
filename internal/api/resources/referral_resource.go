package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type ReferralDetailResource struct {
	ReferralDetails models.ReferralDetails
	moneyConvert    utils.MoneyConvert
}

func (r *ReferralDetailResource) ToJSON() (map[string]interface{}, error) {
	var details map[string]interface{}

	details = map[string]interface{}{
		"referral_code":  r.ReferralDetails.ReferralCode,
		"total_earnings": r.moneyConvert.FromCentsToVault(r.ReferralDetails.TotalEarnings),
		"referred_users": r.ReferralDetails.ReferredUsers,
	}

	return details, nil
}
