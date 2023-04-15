package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/request"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"gorm.io/gorm"
	"net/http"
)

type ReferralController struct {
	userService     services.UserService
	referralService services.ReferralService
	errorHandler    utils.Errors
}

func (rc ReferralController) StoreCode(c *gin.Context) {
	var err error

	user, err := rc.userService.AuthUser(c)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	var store request.StoreUserReferralCode

	if err = c.ShouldBindJSON(&store); err != nil {
		rc.errorHandler.HandleError(c, http.StatusBadRequest, "Error binding JSON", err)
		return
	}

	_, err = rc.referralService.StoreReferralCode(&user, &store)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusInternalServerError, "Error storing referral code", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Referral code stored"})
}

func (rc ReferralController) Details(c *gin.Context) {
	var err error

	user, err := rc.userService.AuthUser(c)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if user.ReferralCode == nil {
		rc.errorHandler.HandleError(c, http.StatusNotFound, "Referral code not found", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"referral_code": user.ReferralCode})
}

func GetUserReferralDetails(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("id")
		var user models.User
		if err := db.Preload("ReferralTier").Preload("ReferralTransactions").Preload("ReferredUsers").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		referralTiers := getReferralTiers(db)
		currentTierCommission := 0.0
		if user.ReferralTierLevel > 0 {
			for _, tier := range referralTiers {
				if tier.Level == int(user.ReferralTierLevel) {
					currentTierCommission = tier.BonusPercentage
					break
				}
			}
			totalEarnings := 0.0
			for _, transaction := range user.ReferralTransactions {
				totalEarnings += transaction.Amount
			}
			responseData := map[string]interface{}{
				"referral_code":           user.ReferralCode,
				"total_earnings":          totalEarnings,
				"current_tier_commission": currentTierCommission,
				"referred_users":          user.ReferredUsers,
			}

			c.JSON(http.StatusOK, responseData)
		}
	}
}

func getReferralTiers(db *gorm.DB) []models.ReferralTier {
	var referralTiers []models.ReferralTier
	if err := db.Find(&referralTiers).Error; err != nil {
		return []models.ReferralTier{}
	}
	return referralTiers
}
