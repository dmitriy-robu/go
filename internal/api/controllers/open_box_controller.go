package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
)

type OpenBoxController struct {
	openBoxService services.OpenBoxManager
}

func NewOpenBoxController(
	openBoxManager services.OpenBoxManager,
) OpenBoxController {
	return OpenBoxController{
		openBoxService: openBoxManager,
	}
}

func (obc OpenBoxController) Open(c *gin.Context) {
	var (
		response   map[string]interface{}
		err        *utils.Errors
		winItem    models.BoxItem
		serverSeed string
	)

	winItem, serverSeed, err = obc.openBoxService.Open(c)

	if err != nil {
	}

	resource := resources.OpenBoxItemResource{
		Item:       winItem,
		ServerSeed: serverSeed,
	}

	response, err = resource.ToJSON()

	if err != nil {
	}

	c.JSON(200, response)
}
