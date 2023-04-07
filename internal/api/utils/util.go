package utils

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HandleError(c *gin.Context, httpStatus int, errMsg string, err error) {
	if err != nil {
		c.JSON(httpStatus, gin.H{
			"error": errMsg,
		})
		log.Println(err)
	}
}
