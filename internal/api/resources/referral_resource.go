package resources

import (
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/utils"
)

type ReferralDetailResource struct {
	ReferralDetails *models.ReferralDetails
	moneyConvert    utils.MoneyConvert
}

func (r *ReferralDetailResource) ToJSON() (map[string]interface{}, error) {
	var (
		err     error
		res     ReferralUserResource
		details map[string]interface{}
	)

	res = ReferralUserResource{
		User: &r.ReferralDetails.ReferredUsers,
	}

	referredUsersJSON, err := res.ToJSON()
	if err != nil {
		return nil, err
	}

	details = map[string]interface{}{
		"referral_code":  r.ReferralDetails.ReferralCode,
		"total_earnings": r.moneyConvert.FromCentsToVault(r.ReferralDetails.TotalEarnings),
		"referred_users": referredUsersJSON,
	}

	return details, nil
}
