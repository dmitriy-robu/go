package services

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/repositories"
	"net/http"
)

type SteamAuthService struct {
	userService     UserService
	steamRepository repositories.SteamRepository
}

func (sam SteamAuthService) setProvider() {
	gothic.GetProviderName = func(*http.Request) (string, error) {
		return "steam", nil
	}
}

func (sam SteamAuthService) Login(c *gin.Context) {
	sam.setProvider()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (sam SteamAuthService) Callback(c *gin.Context) error {
	sam.setProvider()

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return errors.Wrap(err, "Error completing user auth")
	}

	userUuid, err := sam.userService.CreateOrUpdateSteamUser(user)
	if err != nil {
		return errors.Wrap(err, "Error creating or updating user")
	}

	session := sessions.Default(c)
	session.Set("userUuid", userUuid)
	err = session.Save()
	if err != nil {
		return errors.Wrap(err, "Error saving session")
	}

	return nil
}
