package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth"
	"github.com/pkg/errors"
	"go-rust-drop/internal/api/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

type PlayerSummariesResponse struct {
	Response struct {
		Players []models.UserSteamProfile `json:"players"`
	} `json:"response"`
}

type SteamAuthManager struct{}

func (sam SteamAuthManager) Login(c *gin.Context) (string, error) {
	provider, err := goth.GetProvider("steam")
	if err != nil {
		return "", errors.Wrap(err, "Error getting steam provider")
	}

	session, err := provider.BeginAuth("")
	if err != nil {
		return "", errors.Wrap(err, "Error starting steam auth")
	}

	sess := sessions.Default(c)
	sess.Set("goth_session", session.Marshal())
	if err := sess.Save(); err != nil {
		return "", errors.Wrap(err, "Error saving session")

	}

	authURL, err := session.GetAuthURL()
	if err != nil {
		return "", errors.Wrap(err, "Error getting steam auth URL")
	}

	return authURL, nil
}

func (sam SteamAuthManager) Callback(c *gin.Context) error {
	_, err := goth.GetProvider("steam")
	if err != nil {
		return errors.Wrap(err, "Error getting steam provider")
	}

	claimedID, err := url.Parse(c.Request.URL.Query().Get("openid.claimed_id"))
	if err != nil {
		return errors.Wrap(err, "Error parsing claimed ID")
	}
	steamID := path.Base(claimedID.Path)

	userInfo, err := sam.fetchSteamUserInfo(steamID)
	if err != nil {
		return errors.Wrap(err, "Error fetching user info from Steam API")
	}

	userID, err := UserService{}.CreateSteamUser(userInfo)
	if err != nil {
		return errors.Wrap(err, "Error saving user info to MySQL")
	}

	token, err := sam.generateJWTToken(userID)
	if err != nil {
		return errors.Wrap(err, "Error generating JWT token")

	}

	err = UserService{}.InsertUserAuthSteam(userID, userInfo.SteamID)
	if err != nil {
		return errors.Wrap(err, "Error saving user to database")
	}

	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(c.Writer, cookie)

	return nil
}

func (sam SteamAuthManager) fetchSteamUserInfo(steamID string) (models.UserSteamInfo, error) {
	userProfile, err := sam.getSteamUserProfile(steamID)
	if err != nil {
		return models.UserSteamInfo{}, errors.Wrap(err, "Error fetching user profile from Steam API")
	}

	userInfo := models.UserSteamInfo{
		SteamID:   userProfile.SteamID,
		AvatarURL: userProfile.AvatarURL,
		Name:      userProfile.Name,
	}

	return userInfo, nil
}

func (sam SteamAuthManager) getSteamUserProfile(steamID string) (models.UserSteamProfile, error) {
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

func (sam SteamAuthManager) generateJWTToken(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", fmt.Errorf("jwt secret key is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing jwt token: %v", err)
	}

	return signedToken, nil
}
