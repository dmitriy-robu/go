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
	errorHandler utils.Errors
	BoxService   services.BoxService
}

func (b BoxController) Index(c *gin.Context) {
	var (
		boxes         models.Boxes
		boxesResource []map[string]interface{}
	)

	boxes = b.BoxService.FindAll()

	resource := resources.BoxesResource{
		Boxes: boxes,
	}

	boxesResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxesResource)
}

func (b BoxController) Show(c *gin.Context) {
	var (
		err         error
		box         models.Box
		boxResource map[string]interface{}
	)

	box, err = b.BoxService.FindByUUID(c.Param("uuid"))

	if err != nil {
		b.errorHandler.HandleError(c, 404, err.Error(), err)
		return
	}

	resource := resources.BoxResource{
		Box: box,
	}

	boxResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxResource)
}
