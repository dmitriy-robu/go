package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"gorm.io/gorm"
	"net/http"
)

type UserController struct {
	userService  services.UserService
	db           *gorm.DB
	errorHandler utils.Errors
}

func (u UserController) UserInfo(c *gin.Context) {
	var err error

	userID, err := getUserIDFromContext(c)
	if err != nil {
		u.errorHandler.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	userWithBalance, err := u.userService.GetUserInfo(userID)
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

func getUserIDFromContext(c *gin.Context) (string, error) {
	steamUserIDValue, ok := c.MustGet("user_uuid").(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in context")
	}

	return steamUserIDValue, nil
}
