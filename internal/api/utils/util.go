package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Errors struct {
}

func (e Errors) HandleError(c *gin.Context, httpStatus int, errMsg string, err error) {
	if err != nil {
		c.JSON(httpStatus, gin.H{
			"error": errMsg,
		})
		log.Println(err)
	}
}
