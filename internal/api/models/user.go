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
	Name                 string       `gorm:"type:varchar(255)"`
	AvatarURL            string       `gorm:"column:avatar_url"`
	Email                string       `gorm:"type:varchar(255);default:null"`
	EmailVerifiedAt      sql.NullTime `gorm:"column:email_verified_at"`
	Password             string       `gorm:"type:varchar(255);default:null"`
	SteamTradeURL        string       `gorm:"column:steam_trade_url;default:null"`
	Experience           uint64       `gorm:"column:experience;default:0"`
	Active               bool         `gorm:"column:active"`
	IsBot                bool         `gorm:"column:is_bot"`
	ReferralCode         string       `gorm:"column:referral_code;default:null" unique:"true"`
	ReferralTierLevel    uint8        `gorm:"column:referral_tier_level;default:0"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	UserBalance          UserBalance          `gorm:"foreignKey:UserID"`
	ReferralTransactions ReferralTransactions `gorm:"many2many:referral_transactions_users"`
	ReferralUsers        Users                `gorm:"many2many:referrals;joinForeignKey:ReferralUserID;JoinReferences:ParentUserID"`
}

type UserInventory struct {
	Assets              []interface{}
	TotalInventoryCount int
}
