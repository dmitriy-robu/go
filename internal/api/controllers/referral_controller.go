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
	userManager     services.UserManager
	referralManager services.ReferralManager
}

func NewReferralController(
	userManager services.UserManager,
	referralManager services.ReferralManager,
) ReferralController {
	return ReferralController{
		userManager:     userManager,
		referralManager: referralManager,
	}
}

func (rc ReferralController) StoreCode(c *gin.Context) {
	var (
		err          error
		user         models.User
		store        requests.StoreUserReferralCode
		errorHandler *utils.Errors
	)

	user, errorHandler = rc.userManager.AuthUser(c)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	if err = c.ShouldBindJSON(&store); err != nil {
		errorHandler = utils.NewErrors(http.StatusBadRequest, "Error binding JSON", err)
		errorHandler.HandleError(c)
		return
	}

	_, errorHandler = rc.referralManager.StoreReferralCode(user, store)
	if errorHandler != nil {
		errorHandler.HandleError(c)
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
		errorHandler            *utils.Errors
	)

	user, errorHandler = rc.userManager.AuthUser(c)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	referralDetails, errorHandler = rc.referralManager.GetReferralDetails(user)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	referralDetailResources = resources.ReferralDetailResource{
		ReferralDetails: referralDetails,
	}

	referralDetailResource, err = referralDetailResources.ToJSON()
	if err != nil {
		errorHandler = utils.NewErrors(http.StatusInternalServerError, "Error converting to JSON", err)
		errorHandler.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, referralDetailResource)
}
