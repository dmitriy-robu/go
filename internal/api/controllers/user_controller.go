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
}

func (u UserController) UserInfo(c *gin.Context) {
	var (
		err           error
		user          models.User
		userInfo      map[string]interface{}
		userResources resources.UserResources
		errorHandler  utils.Errors
	)

	user, errorHandler = u.userManager.AuthUser(c)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	user, errorHandler = u.userManager.GetUserWithBalance(user)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	userResources = resources.UserResources{
		User: &user,
	}

	userInfo, err = userResources.ToJSON()
	if err != nil {
		errorHandler = utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error converting user information to JSON",
			Err:     err,
		}
		errorHandler.HandleError(c)
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
		errorHandler           utils.Errors
	)

	user, errorHandler = u.userManager.AuthUser(c)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	inventory, errorHandler = u.userInventoryManager.GetInventoryForUser(user.UUID)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	userInventoryResources = resources.UserInventoryResources{
		AssetData: inventory.AssetData,
	}

	userInventoryResource, err = userInventoryResources.ToJSON()
	if err != nil {
		errorHandler = utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error converting user information to JSON",
			Err:     err,
		}
		errorHandler.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, userInventoryResource)
}

func (u UserController) StoreSteamTradeURL(c *gin.Context) {
	var (
		err          error
		store        requests.StoreUserSteamTradeURL
		user         models.User
		errorHandler utils.Errors
	)

	user, errorHandler = u.userManager.AuthUser(c)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	if err = c.BindJSON(&store); err != nil {
		errorHandler = utils.Errors{
			Code:    http.StatusBadRequest,
			Message: "Error binding trade URL",
			Err:     err,
		}
		errorHandler.HandleError(c)
		return
	}

	if errorHandler = u.userManager.StoreSteamTradeURL(user, store); errorHandler.Err != nil {
		errorHandler.HandleError(c)
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
		errorHandler                utils.Errors
	)

	user, errorHandler = u.userManager.AuthUser(c)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	userWithBalance, errorHandler = u.userManager.GetUserWithBalance(user)
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	userLevel = u.levelManager.GetLevelForByExperience(*user.Experience)

	userUpdatableFieldsResource = resources.UserUpdatableFieldsResource{
		User:  userWithBalance,
		Level: userLevel,
	}

	userInfo, err = userUpdatableFieldsResource.ToJSON()
	if err != nil {
		errorHandler = utils.Errors{
			Code:    http.StatusInternalServerError,
			Message: "Error converting user information to JSON",
			Err:     err,
		}
		errorHandler.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
