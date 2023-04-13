package services

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"net/http"
)

type SteamAuthService struct {
	userService     UserService
	steamRepository repositories.SteamRepository
}

type PlayerSummariesResponse struct {
	Response struct {
		Players []models.UserSteamProfile `json:"players"`
	} `json:"response"`
}

func setProvider() {
	gothic.GetProviderName = func(*http.Request) (string, error) {
		return "steam", nil
	}
}

func (sam SteamAuthService) Login(c *gin.Context) {
	setProvider()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (sam SteamAuthService) Callback(c *gin.Context) error {
	setProvider()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return errors.Wrap(err, "Error completing user auth")
	}

	session := sessions.Default(c)
	session.Set("steamID", user.UserID)
	err = session.Save()
	if err != nil {
		return errors.Wrap(err, "Error saving session")
	}

	if err = sam.userService.CreateOrUpdateSteamUser(user); err != nil {
		return errors.Wrap(err, "Error creating or updating user")
	}

	return nil
}
