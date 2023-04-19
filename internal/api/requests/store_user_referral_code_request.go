package requests

type StoreUserReferralCode struct {
	ReferralCode string `json:"referral_code" binding:"required,max=255"`
}
