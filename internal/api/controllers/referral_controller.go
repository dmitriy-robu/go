package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/mappers"
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

	var store mappers.StoreUserReferralCode

	if err = c.ShouldBindJSON(&store); err != nil {
		rc.errorHandler.HandleError(c, http.StatusBadRequest, "Error binding JSON", err)
		return
	}

	_, err = rc.referralService.StoreReferralCode(&user, &store)

}
