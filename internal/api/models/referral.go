package models

import "time"

type ReferralTier struct {
	ID                  uint    `gorm:"primaryKey"`
	Level               int     `gorm:"column:level"`
	RequiredSpentAmount float64 `gorm:"column:required_spent_amount"`
	BonusPercentage     float64 `gorm:"column:bonus_percentage"`
	RecommendedPersons  int     `gorm:"column:recommended_persons"`
}

type ReferralTransaction struct {
	ID         uint `gorm:"primaryKey"`
	ReferralID uint `gorm:"column:referral_id"`
	Amount     int  `gorm:"column:amount"`
}

type Referral struct {
	ID             uint    `gorm:"primaryKey"`
	ParentUserID   uint    `gorm:"column:parent_user_id"`
	ReferralUserID uint    `gorm:"column:referral_user_id"`
	ReceivedAmount float64 `gorm:"column:received_amount"`
}

type ReferralDetails struct {
	ReferralCode          *string        `json:"referral_code"`
	TotalEarnings         int            `json:"total_earnings"`
	CurrentTierCommission float64        `json:"current_tier_commission"`
	ReferredUsers         []ReferredUser `json:"referred_users"`
}

type ReferredUser struct {
	Name              string
	TotalEarnings     int
	EarningCommission float64
	CurrentTier       uint
	CreatedAt         time.Time
}
