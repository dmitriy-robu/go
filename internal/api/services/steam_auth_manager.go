package services

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/repositories"
	"net/http"
)

type SteamAuthManager struct {
	userManager     UserManager
	steamRepository repositories.SteamRepository
}

func (sam SteamAuthManager) setProvider() {
	gothic.GetProviderName = func(*http.Request) (string, error) {
		return "steam", nil
	}
}

func (sam SteamAuthManager) Login(c *gin.Context) {
	sam.setProvider()

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (sam SteamAuthManager) Callback(c *gin.Context) error {
	var (
		err      error
		user     goth.User
		userUuid string
		session  sessions.Session
	)

	sam.setProvider()

	user, err = gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return errors.Wrap(err, "Error completing user auth")
	}

	userUuid, err = sam.userManager.CreateOrUpdateSteamUser(user)
	if err != nil {
		return errors.Wrap(err, "Error creating or updating user")
	}

	session = sessions.Default(c)
	session.Set("userUuid", userUuid)
	if err = session.Save(); err != nil {
		return errors.Wrap(err, "Error saving session")
	}

	return nil
}
