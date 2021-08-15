package Product_routes

import (
	"practice1servergo/Controllers"

	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get all products",
	})
}

func Routes(route *gin.Engine) {
	productRoutes := route.Group("/product")
	productRoutes.GET("", getAll)
	productRoutes.POST("", Controllers.AddNewProduct)
}
