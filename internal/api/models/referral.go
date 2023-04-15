package models

type ReferralTier struct {
	ID                  uint    `gorm:"primaryKey"`
	Level               int     `gorm:"column:level"`
	RequiredSpentAmount float64 `gorm:"column:required_spent_amount"`
	BonusPercentage     float64 `gorm:"column:bonus_percentage"`
	RecommendedPersons  int     `gorm:"column:recommended_persons"`
}

type ReferralTransaction struct {
	ID         uint    `gorm:"primaryKey"`
	ReferralID uint    `gorm:"column:referral_id"`
	Amount     float64 `gorm:"column:amount"`
}

type Referral struct {
	ID             uint    `gorm:"primaryKey"`
	ParentUserID   uint    `gorm:"column:parent_user_id"`
	ReferralUserID uint    `gorm:"column:referral_user_id"`
	ReceivedAmount float64 `gorm:"column:received_amount"`
}
