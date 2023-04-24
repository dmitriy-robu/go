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
	steamAuthManager services.SteamAuthManager
}

func NewSteamAuthController(
	steamAuthManager services.SteamAuthManager,
) SteamAuthController {
	return SteamAuthController{
		steamAuthManager: steamAuthManager,
	}
}

func (s SteamAuthController) Login(c *gin.Context) {
	s.steamAuthManager.Login(c)
}

func (s SteamAuthController) Callback(c *gin.Context) {
	var (
		errorHandler *utils.Errors
	)

	errorHandler = s.steamAuthManager.Callback(c)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")
}
