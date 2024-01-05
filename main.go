package main

import (
	"github.com/caio-gferreira/BUBBLE-STATIONET-SHOP/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	application := gin.Default()

	routes.Initialize()

	application.Run(":8000")
}
