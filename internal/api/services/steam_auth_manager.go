package services

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"go-rust-drop/internal/api/repositories"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type SteamAuthManager struct {
	userManager     UserManager
	steamRepository repositories.SteamRepository
}

func NewSteamAuthManager(
	userManager UserManager,
	steamRepository repositories.SteamRepository,
) SteamAuthManager {
	return SteamAuthManager{
		userManager:     userManager,
		steamRepository: steamRepository,
	}
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

func (sam SteamAuthManager) Callback(c *gin.Context) *utils.Errors {
	var (
		err          error
		user         goth.User
		userUuid     string
		session      sessions.Session
		errorHandler *utils.Errors
	)

	sam.setProvider()

	user, err = gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {

		return utils.NewErrors(http.StatusInternalServerError, "Error completing user auth", err)
	}

	userUuid, errorHandler = sam.userManager.CreateOrUpdateSteamUser(user)
	if errorHandler != nil {
		return errorHandler
	}

	session = sessions.Default(c)
	session.Set("userUuid", userUuid)
	if err = session.Save(); err != nil {
		return utils.NewErrors(http.StatusInternalServerError, "Error saving session", err)
	}

	return nil
}
