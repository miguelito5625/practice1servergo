package Systemroutes

import (
	"practice1servergo/Controllers"
	"practice1servergo/Middlewares"

	"github.com/gin-gonic/gin"
)

func Rol_Routes(route *gin.Engine) {
	rolRoutes := route.Group("/rol")
	rolRoutes.Use(Middlewares.CheckAuth())
	{

		rolRoutes.GET("", Controllers.ListRol)
		rolRoutes.GET("/:id", Controllers.GetOneRol)
		rolRoutes.POST("", Controllers.AddNewRol)
		rolRoutes.PUT("/:id", Controllers.PutOneRol)
		rolRoutes.DELETE("/:id", Controllers.DeleteRol)

	}

}
