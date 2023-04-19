package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"log"
	"net/http"
)

type UserController struct {
	userService          services.UserService
	userInventoryService services.UserInventoryService
	errorHandler         utils.Errors
}

func (u UserController) UserInfo(c *gin.Context) {
	var err error

	user, err := u.userService.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	user, err = u.userService.GetUserWithBalance(user)

	userResources := resources.UserResources{
		User: &user,
	}

	userInfo, err := userResources.ToJSON()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}

func (u UserController) UserInventory(c *gin.Context) {
	var err error

	user, err := u.userService.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	inventory, err := u.userInventoryService.GetInventoryForUser(&user.UUID)
	log.Printf("inventory: %+v", inventory)
	if err != nil || inventory == nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting user inventory", err)
		return
	}

	userResources := resources.UserInventoryResources{
		AssetData: inventory.AssetData,
	}

	userInventoryResource, err := userResources.ToJSON()
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

	user, err = u.userService.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	if err = c.BindJSON(&store); err != nil {
		u.errorHandler.HandleError(c, http.StatusBadRequest, "Error binding trade URL", err)
		return
	}

	if err = u.userService.StoreSteamTradeURL(user, store); err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error updating user trade URL", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Trade URL updated",
	})
}
