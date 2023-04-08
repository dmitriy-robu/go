package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type SteamAuthController struct {
	steamAuthManager services.SteamAuthManager
}

func (s SteamAuthController) Login(c *gin.Context) {
	authURL, err := s.steamAuthManager.Login(c)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (s SteamAuthController) Callback(c *gin.Context) {
	err := s.steamAuthManager.Callback(c)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
