package routes

import "github.com/gin-gonic/gin"

func SetRoutes(r *gin.Engine) {

	api := r.Group("/api/v1")
	{
		api.GET("health", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"status": "success",
			})
		})
	}
}
