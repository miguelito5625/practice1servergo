package Product_routes

import (
	"practice1servergo/Controllers"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	productRoutes := route.Group("/product")
	productRoutes.GET("", Controllers.ListProduct)
	productRoutes.POST("", Controllers.AddNewProduct)
}
