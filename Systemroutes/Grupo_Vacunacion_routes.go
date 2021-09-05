package Systemroutes

import (
	"practice1servergo/Controllers"
	"practice1servergo/Middlewares"

	"github.com/gin-gonic/gin"
)

func Grupo_Vacunacion_Routes(route *gin.Engine) {
	grupo_vacunacionRoutes := route.Group("/grupo_vacunacion")
	grupo_vacunacionRoutes.Use(Middlewares.CheckAuth())
	{

		grupo_vacunacionRoutes.GET("", Controllers.ListGrupo_Vacunacion)
		grupo_vacunacionRoutes.GET("/:id", Controllers.GetOneGrupo_Vacunacion)
		grupo_vacunacionRoutes.POST("", Controllers.AddNewGrupo_Vacunacion)
		grupo_vacunacionRoutes.PUT("/:id", Controllers.PutOneGrupo_Vacunacion)
		grupo_vacunacionRoutes.DELETE("/:id", Controllers.DeleteGrupo_Vacunacion)

	}

}
