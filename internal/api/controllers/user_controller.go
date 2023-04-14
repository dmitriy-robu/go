package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type UserController struct {
	userService  services.UserService
	errorHandler utils.Errors
}

func (u UserController) UserInfo(c *gin.Context) {
	var err error

	user, err := u.userService.AuthUser(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	userWithBalance, err := u.userService.GetUserInfo(user)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error getting user information", err)
		return
	}

	userResources := resources.UserResources{
		UserBalance: userWithBalance.UserBalance,
		User:        userWithBalance.User,
	}

	jsonData, err := userResources.UserInfo()
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}
