package routes

import (
	"net/http"

	"github.com/caio-gferreira/BUBBLE-STATIONET-SHOP/controllers"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	Router := gin.Default()

	Router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	Router.GET("/product", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": controllers.GetProduct(),
		})
	})
}
