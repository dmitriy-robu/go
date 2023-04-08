package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type SteamAuthController struct {
	steamAuthService services.SteamAuthService
}

func (s SteamAuthController) Login(c *gin.Context) {
	authURL, err := s.steamAuthService.Login(c)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (s SteamAuthController) Callback(c *gin.Context) {
	err := s.steamAuthService.Callback(c)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
