package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/request"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
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
