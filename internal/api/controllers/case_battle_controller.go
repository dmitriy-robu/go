package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/requests"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type CaseBattleController struct {
	userManager       services.UserManager
	caseBattleManager services.CaseBattleManager
}

func (cbc CaseBattleController) Create(c *gin.Context) {
	var (
		err               error
		user              models.User
		caseBattleRequest requests.CaseBattleStoreRequest
		boxUUID           string
		errorHandler      *utils.Errors
	)

	if err = c.ShouldBindJSON(&caseBattleRequest); err != nil {
		errorHandler = utils.NewErrors(http.StatusBadRequest, "Bad request", err)
		errorHandler.HandleError(c)
		return
	}

	if err = validator.New().Struct(caseBattleRequest); err != nil {
		errorHandler = utils.NewErrors(http.StatusBadRequest, "Bad request", err)
		errorHandler.HandleError(c)
		return
	}

	user, errorHandler = cbc.userManager.AuthUser(c)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	boxUUID, errorHandler = cbc.caseBattleManager.Create(caseBattleRequest, user)
	if errorHandler != nil {
		errorHandler.HandleError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"box_uuid": boxUUID})
}

func (cbc CaseBattleController) Join(c *gin.Context) {
	// 1. Получить UUID игры из параметров URL
	// 2. Найти игру по UUID
	// 3. Проверить, есть ли свободные места для присоединения к игре
	// 4. Если есть место, добавить игрока и сохранить изменения в базе данных или кэше
	// 5. Отправить ответ с обновленной информацией об игре
}

func (cbc CaseBattleController) Start(c *gin.Context) {
	// 1. Получить UUID игры из параметров URL
	// 2. Найти игру по UUID
	// 3. Если игроки присоединились к игре, начать игру, иначе добавить ботов
	// 4. Выполнить игровую логику в зависимости от выбранного режима (Group, Standard, Crazy)
	// 5. Обновить статус игры, распределить призы между игроками и сохранить результаты в базе данных или кэше.
	// 6. Отправить ответ с результатами игры и информацией о выигранных предметах.
}
