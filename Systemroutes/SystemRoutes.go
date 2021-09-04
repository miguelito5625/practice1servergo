package Systemroutes

import (
	"practice1servergo/Systemroutes/Rol_routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	Rol_routes.Routes(r)

	return r
}
