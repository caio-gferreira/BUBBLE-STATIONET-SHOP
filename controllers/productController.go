package controllers

import (
	"github.com/caio-gferreira/BUBBLE-STATIONET-SHOP/models"
)

func GetProduct() []models.Product {
	return []models.Product{
		{Sku: 12, Amount: 3, Price_unit: 33.23, Product_Name: "Caneta"},
	}

	// var ctx *gin.Context
	// ctx.JSON(200, gin.H{
	// 	"product": product,
	// })
}
