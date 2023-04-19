package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/steam"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
	"os"
)

func init() {
	goth.UseProviders(
		steam.New(os.Getenv("STEAM_API_KEY"), os.Getenv("STEAM_CALLBACK_URL")),
	)
}

type SteamAuthController struct {
	steamAuthService services.SteamAuthService
	errorHandler     utils.Errors
}

func (s SteamAuthController) Login(c *gin.Context) {
	s.steamAuthService.Login(c)
}

func (s SteamAuthController) Callback(c *gin.Context) {
	var err error

	err = s.steamAuthService.Callback(c)
	if err != nil {
		s.errorHandler.HandleError(c, http.StatusInternalServerError, "Error logging in", err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
