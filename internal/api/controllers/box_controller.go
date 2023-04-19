package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rust-drop/internal/api/models"
	"go-rust-drop/internal/api/resources"
	"go-rust-drop/internal/api/services"
	"net/http"
)

type BoxController struct {
	BoxService services.BoxService
}

func (b BoxController) Index(c *gin.Context) {
	var (
		boxes         models.Boxes
		boxesResource []map[string]interface{}
		err           error
	)

	boxes = b.BoxService.FindAll()

	resource := resources.BoxesResource{
		Boxes: boxes,
	}

	boxesResource, err = resource.ToJSON()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting boxes"})
		return
	}

	c.JSON(http.StatusOK, boxesResource)
}

func (b BoxController) Show(c *gin.Context) {
	var (
		//box         models.Box
		boxResource map[string]interface{}
		err         error
	)

	b.BoxService.FindByUUID(c.Param("uuid"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting box"})
		return
	}

	c.JSON(http.StatusOK, boxResource)
}
