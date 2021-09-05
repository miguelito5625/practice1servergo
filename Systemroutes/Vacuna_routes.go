package Systemroutes

import (
	"practice1servergo/Controllers"
	"practice1servergo/Middlewares"

	"github.com/gin-gonic/gin"
)

func Vacuna_Routes(route *gin.Engine) {
	vacunaRoutes := route.Group("/vacuna")
	vacunaRoutes.Use(Middlewares.CheckAuth())
	{

		vacunaRoutes.GET("", Controllers.ListVacuna)
		vacunaRoutes.GET("/:id", Controllers.GetOneVacuna)
		vacunaRoutes.POST("", Controllers.AddNewVacuna)
		vacunaRoutes.PUT("/:id", Controllers.PutOneVacuna)
		vacunaRoutes.DELETE("/:id", Controllers.DeleteVacuna)

	}

}
