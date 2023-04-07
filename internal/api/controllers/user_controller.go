package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type UserController struct {
}

func (u UserController) UserInfo(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		utils.HandleError(c, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	userWithBalance, err := services.UserService{}.GetUserInfo(userID)

	userResources := resources.UserResources{
		UserBalance: userWithBalance.UserBalance,
		User:        userWithBalance.User,
	}

	jsonData, err := userResources.UserInfo()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error converting user information to JSON", err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}

func getUserIDFromContext(c *gin.Context) (uint64, error) {
	userIDValue, ok := c.Get("userID")
	if !ok {
		return 0, fmt.Errorf("user ID not found in context")
	}

	userID, ok := userIDValue.(uint64)
	if !ok {
		return 0, fmt.Errorf("user ID has wrong type")
	}

	return userID, nil
}
