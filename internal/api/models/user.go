package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Users []User

// User полная таблица в mysql
type User struct {
	gorm.Model
	ID                   uint         `gorm:"primaryKey"`
	UUID                 string       `gorm:"unique"`
	Name                 *string      `gorm:"type:varchar(255)"`
	AvatarURL            *string      `gorm:"column:avatar_url"`
	Email                *string      `gorm:"type:varchar(255)"`
	EmailVerifiedAt      sql.NullTime `gorm:"column:email_verified_at"`
	Password             *string      `gorm:"type:varchar(255)"`
	SteamTradeURL        *string      `gorm:"column:steam_trade_url"`
	Experience           *int         `gorm:"column:experience,default:0"`
	Active               bool         `gorm:"column:active"`
	IsBot                bool         `gorm:"column:is_bot"`
	RememberToken        *string      `gorm:"column:remember_token"`
	ReferralCode         *string      `gorm:"column:referral_code" unique:"true"`
	ReferralTierLevel    uint         `gorm:"column:referral_tier_level"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	UserBalance          UserBalance           `gorm:"foreignKey:UserID"`
	ReferralTier         ReferralTier          `gorm:"foreignKey:ReferralTierLevel"`
	ReferralTransactions []ReferralTransaction `gorm:"many2many:referral_transactions_users"`
	ReferralUsers        []User                `gorm:"many2many:referrals;joinForeignKey:ReferralUserID;JoinReferences:ParentUserID"`
}

type UserInventory struct {
	Assets              []interface{}
	TotalInventoryCount int
}
