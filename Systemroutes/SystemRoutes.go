package Systemroutes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	Rol_Routes(r)
	Usuario_Routes(r)
	Vacuna_Routes(r)
	Centro_Vacunacion_Routes(r)
	Grupo_Vacunacion_Routes(r)

	return r
}
