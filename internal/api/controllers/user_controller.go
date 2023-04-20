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

type UserController struct {
	userManager          services.UserManager
	levelManager         services.LevelManager
	userInventoryManager services.UserInventoryManager
	errorHandler         utils.Errors
}

func (u UserController) UserInfo(c *gin.Context) {
	var (
		err           error
		user          models.User
		userInfo      map[string]interface{}
		userResources resources.UserResources
	)

	user, err = u.userManager.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	user, err = u.userManager.GetUserWithBalance(user)

	userResources = resources.UserResources{
		User: &user,
	}

	userInfo, err = userResources.ToJSON()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (u UserController) UserInventory(c *gin.Context) {
	var (
		err                    error
		user                   models.User
		inventory              models.InventoryData
		userInventoryResources resources.UserInventoryResources
		userInventoryResource  []map[string]interface{}
	)

	user, err = u.userManager.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	inventory, err = u.userInventoryManager.GetInventoryForUser(user.UUID)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting user inventory", err)
		return
	}

	userInventoryResources = resources.UserInventoryResources{
		AssetData: inventory.AssetData,
	}

	userInventoryResource, err = userInventoryResources.ToJSON()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.JSON(http.StatusOK, userInventoryResource)
}

func (u UserController) StoreSteamTradeURL(c *gin.Context) {
	var (
		err   error
		store requests.StoreUserSteamTradeURL
		user  models.User
	)

	user, err = u.userManager.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if err = c.BindJSON(&store); err != nil {
		u.errorHandler.HandleError(c, http.StatusBadRequest, "Error binding trade URL", err)
		return
	}

	if err = u.userManager.StoreSteamTradeURL(user, store); err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error updating user trade URL", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Trade URL updated",
	})
}

func (u UserController) GetUpdatableFields(c *gin.Context) {
	var (
		err                         error
		user                        models.User
		userInfo                    map[string]interface{}
		userWithBalance             models.User
		userUpdatableFieldsResource resources.UserUpdatableFieldsResource
		userLevel                   models.Level
	)

	user, err = u.userManager.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	userWithBalance, err = u.userManager.GetUserWithBalance(user)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting user balance", err)
		return
	}

	userLevel = u.levelManager.GetLevelForByExperience(*user.Experience)

	userUpdatableFieldsResource = resources.UserUpdatableFieldsResource{
		User:  userWithBalance,
		Level: userLevel,
	}

	userInfo, err = userUpdatableFieldsResource.ToJSON()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
