package controllers

import (
	"github.com/gin-gonic/gin"
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
		UserBalance: user.UserBalance,
		User:        &user,
	}

	jsonData, err := userResources.UserInfo()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}

func (u UserController) UserInventory(c *gin.Context) {
	var err error

	user, err := u.userService.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	inventory, err := u.userInventoryService.GetInventoryForUser(&user.UUID)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting user inventory", err)
		return
	}

	userResources := resources.UserResources{
		AssetData: inventory.AssetData,
	}
	log.Printf("inventory: %+v", inventory)

	jsonData, err := userResources.UserInventory()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
