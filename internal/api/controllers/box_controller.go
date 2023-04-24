package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"go-rust-drop/internal/api/utils"
	"net/http"
)

type BoxController struct {
	BoxManager services.BoxManager
}

func (b BoxController) Index(c *gin.Context) {
	var (
		boxes         models.Boxes
		boxesResource []map[string]interface{}
	)

	boxes = b.BoxManager.FindAllWithItems()

	resource := resources.BoxesResource{
		Boxes: boxes,
	}

	boxesResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxesResource)
}

func (b BoxController) Show(c *gin.Context) {
	var (
		box          models.Box
		boxResource  map[string]interface{}
		errorHandler *utils.Errors
	)

	box, errorHandler = b.BoxManager.FindByUUIDWithItems(c.Param("uuid"))
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	resource := resources.BoxResource{
		Box: box,
	}

	boxResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxResource)
}

func (b BoxController) Open(c *gin.Context) {
	var (
		box          models.Box
		boxResource  map[string]interface{}
		errorHandler *utils.Errors
	)

	box, errorHandler = b.BoxManager.Open(c.Param("uuid"))
	if errorHandler.Err != nil {
		errorHandler.HandleError(c)
		return
	}

	resource := resources.BoxResource{
		Box: box,
	}

	boxResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxResource)
}
