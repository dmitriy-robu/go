package public

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	get(router)
	post(router)
}

func get(router *gin.RouterGroup) {
	router.GET("/provably-fair", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello",
		})
	})
}

func post(router *gin.RouterGroup) {

}
