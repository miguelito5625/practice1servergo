package Systemroutes

import (
	"practice1servergo/Controllers"
	"practice1servergo/Middlewares"

	"github.com/gin-gonic/gin"
)

func Autenticacion_Routes(route *gin.Engine) {
	autenticacionRoutes := route.Group("/autenticacion")
	autenticacionRoutes.Use(Middlewares.CheckAuth())
	{

		autenticacionRoutes.POST("/login", Controllers.IniciarSesion)

	}

}
