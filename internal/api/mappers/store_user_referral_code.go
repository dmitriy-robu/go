package mappers

type StoreUserReferralCode struct {
	ReferralCode string `json:"referral_code" binding:"required"`
}
