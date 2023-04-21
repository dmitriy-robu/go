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
	BoxManager   services.BoxManager
}

func (b BoxController) Index(c *gin.Context) {
	var (
		boxes         models.Boxes
		boxesResource []map[string]interface{}
	)

	boxes = b.BoxManager.FindAll()

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

	box, err = b.BoxManager.FindByUUID(c.Param("uuid"))

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

func (b BoxController) Open(c *gin.Context) {
	var (
		err         error
		box         models.Box
		boxResource map[string]interface{}
	)

	box, err = b.BoxManager.Open(c.Param("uuid"))

	if err != nil {
		b.errorHandler.HandleError(c, 500, err.Error(), err)
		return
	}

	resource := resources.BoxResource{
		Box: box,
	}

	boxResource = resource.ToJSON()

	c.JSON(http.StatusOK, boxResource)
}
