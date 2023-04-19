package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/resources"
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
	var (
		err   error
		user  models.User
		store requests.StoreUserReferralCode
	)

	user, err = rc.userService.AuthUser(c)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if err = c.ShouldBindJSON(&store); err != nil {
		rc.errorHandler.HandleError(c, http.StatusBadRequest, "Error binding JSON", err)
		return
	}

	_, err = rc.referralService.StoreReferralCode(user, store)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusInternalServerError, "Error storing referral code", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Referral code stored"})
}

func (rc ReferralController) Details(c *gin.Context) {
	var (
		err                     error
		user                    models.User
		referralDetails         models.ReferralDetails
		referralDetailResource  map[string]interface{}
		referralDetailResources resources.ReferralDetailResource
	)

	user, err = rc.userService.AuthUser(c)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	referralDetails, err = rc.referralService.GetReferralDetails(user)
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting referral details", err)
		return
	}

	referralDetailResources = resources.ReferralDetailResource{
		ReferralDetails: referralDetails,
	}

	referralDetailResource, err = referralDetailResources.ToJSON()
	if err != nil {
		rc.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.JSON(http.StatusOK, referralDetailResource)
}
