package Systemroutes

import (
	"practice1servergo/Controllers"
	"practice1servergo/Middlewares"

	"github.com/gin-gonic/gin"
)

func Centro_Vacunacion_Routes(route *gin.Engine) {
	centro_vacunacionRoutes := route.Group("/centro_vacunacion")
	centro_vacunacionRoutes.Use(Middlewares.CheckAuth())
	{

		centro_vacunacionRoutes.GET("", Controllers.ListCentro_Vacunacion)
		centro_vacunacionRoutes.GET("/:id", Controllers.GetOneCentro_Vacunacion)
		centro_vacunacionRoutes.POST("", Controllers.AddNewCentro_Vacunacion)
		centro_vacunacionRoutes.PUT("/:id", Controllers.PutOneCentro_Vacunacion)
		centro_vacunacionRoutes.DELETE("/:id", Controllers.DeleteCentro_Vacunacion)

	}

}
