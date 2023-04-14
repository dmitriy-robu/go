package middlewares_test

import (
	"encoding/json"
	"go-rust-drop/internal/api/middlewares"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Создайте тестовый обработчик, который будет использоваться с middleware
	testHandler := func(c *gin.Context) {
		userID, _ := c.Get("userid")
		c.JSON(http.StatusOK, gin.H{"userid": userID})
	}

	// Настройте middleware и сессии
	Middleware := middlewares.Middleware{}
	store := cookie.NewStore([]byte("test-session-secret"))

	// Создайте Gin Engine и зарегистрируйте middleware и обработчик
	engine := gin.New()
	engine.Use(sessions.Sessions("test-session", store))
	engine.Use(Middleware.VerifyToken)
	engine.GET("/test", testHandler)

	// Создайте тестовый запрос
	req, _ := http.NewRequest("GET", "/test", nil)
	resp := httptest.NewRecorder()

	// Проверьте случай без аутентификации (ожидается 401 Unauthorized)
	engine.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusUnauthorized, resp.Code)

	// Создайте сессию с SteamID (имитируя успешную авторизацию)
	session, _ := store.New(nil, "test-session")
	//session.Set("userid", "test-steamid")
	err := session.Save(req, resp)
	assert.NoError(t, err)

	// Добавьте куки сессии к тестовому запросу
	req.Header.Set("Cookie", resp.Header().Get("Set-Cookie"))

	// Повторите запрос с сессией (ожидается 200 OK)
	resp = httptest.NewRecorder()
	engine.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	var responseJSON map[string]string
	err = json.Unmarshal(resp.Body.Bytes(), &responseJSON)
	assert.NoError(t, err)

	assert.Equal(t, "test-steamid", responseJSON["userid"])
}
