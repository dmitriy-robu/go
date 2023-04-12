package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/repositories"
	"io"
	"net/http"
	"os"
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

func (sam SteamAuthService) Login(c *gin.Context) (string, error) {
	if os.Getenv("SESSION_SECRET") == "" {
		return "", errors.New("SESSION_SECRET env variable is not set")
	}

	provider, err := goth.GetProvider("steam")
	if err != nil {
		return "", errors.Wrap(err, "Error getting steam provider")
	}

	session, err := provider.BeginAuth("oauth_verifier")
	if err != nil {
		return "", errors.Wrap(err, "Error starting steam auth")
	}

	sess := sessions.Default(c)
	sess.Set(os.Getenv("SESSION_SECRET"), session.Marshal())
	if err := sess.Save(); err != nil {
		return "", errors.Wrap(err, "Error saving session")
	}

	authURL, err := session.GetAuthURL()
	if err != nil {
		return "", errors.Wrap(err, "Error getting steam auth URL")
	}

	return authURL, nil
}

func (sam SteamAuthService) Callback(c *gin.Context) error {
	var err error
	providerName := "steam"

	provider, err := goth.GetProvider(providerName)
	if err != nil {
		return errors.Wrap(err, "Error getting steam provider")
	}

	req := c.Request

	value, err := gothic.GetFromSession(providerName, req)
	if err != nil {
		return errors.Wrap(err, "Error getting session")
	}
	// релизазуй
	// defer Logout(res, req)

	session, err := provider.UnmarshalSession(value)
	if err != nil {
		return errors.Wrap(err, "Error unmarshalling session")
	}

	_, err = session.Authorize(provider, c.Request.URL.Query())
	if err != nil {
		return errors.Wrap(err, "Error authorizing session")
	}

	userFetch, err := provider.FetchUser(session)
	if err != nil {
		return errors.Wrap(err, "Error fetching user")
	}

	userSteamInfo := models.UserSteamInfo{
		Name:              &userFetch.Name,
		SteamUserID:       &userFetch.UserID,
		AvatarURL:         &userFetch.AvatarURL,
		AccessToken:       &userFetch.AccessToken,
		AccessTokenSecret: &userFetch.AccessTokenSecret,
		RefreshToken:      &userFetch.RefreshToken,
		ExpiresAt:         &userFetch.ExpiresAt,
	}

	_, err = sam.userService.CreateOrUpdateSteamUser(userSteamInfo)
	if err != nil {
		return errors.Wrap(err, "Error saving user info to database")
	}

	if err = sam.storeUserInSession(userFetch, c); err != nil {
		return errors.Wrap(err, "Error storing user in session")
	}

	c.Redirect(http.StatusTemporaryRedirect, "/")

	return nil
}

func (sam SteamAuthService) storeUserInSession(user goth.User, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set("user", user)
	err := session.Save()
	if err != nil {
		return err
	}

	return nil
}

func (sam SteamAuthService) fetchSteamUserInfo(steamID string) (models.UserSteamInfo, error) {
	userProfile, err := sam.getSteamUserProfile(steamID)
	if err != nil {
		return models.UserSteamInfo{}, errors.Wrap(err, "Error fetching user profile from Steam API")
	}

	userInfo := models.UserSteamInfo{
		SteamUserID: userProfile.SteamID,
		AvatarURL:   userProfile.AvatarURL,
		Name:        userProfile.Name,
	}

	return userInfo, nil
}

func (sam SteamAuthService) getSteamUserProfile(steamID string) (models.UserSteamProfile, error) {
	var steamApiKey = os.Getenv("STEAM_API_KEY")
	if steamApiKey == "" {
		return models.UserSteamProfile{}, fmt.Errorf("steam api key is not set")
	}

	var steamApiUrl = os.Getenv("STEAM_API_URL")
	if steamApiUrl == "" {
		return models.UserSteamProfile{}, fmt.Errorf("steam api url is not set")
	}

	steamUrl := fmt.Sprintf("%s/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s", steamApiUrl, steamApiKey, steamID)

	resp, err := http.Get(steamUrl)
	if err != nil {
		return models.UserSteamProfile{}, fmt.Errorf("error fetching user profile from Steam API: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.UserSteamProfile{}, fmt.Errorf("error reading response body: %v", err)
	}

	var playerSummariesResponse PlayerSummariesResponse
	err = json.Unmarshal(body, &playerSummariesResponse)
	if err != nil {
		return models.UserSteamProfile{}, fmt.Errorf("error unmarshalling player summaries response: %v", err)
	}

	if len(playerSummariesResponse.Response.Players) == 0 {
		return models.UserSteamProfile{}, fmt.Errorf("no players found for steamID: %s", steamID)
	}

	return playerSummariesResponse.Response.Players[0], nil
}
