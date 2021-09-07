package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"

	"github.com/gin-gonic/gin"
)

type Credenciales struct {
	Cui   string `json:"cui"`
	Clave string `json:"clave"`
}

func IniciarSesion(c *gin.Context) {
	var credenciales Credenciales
	c.BindJSON(&credenciales)
	log.Println(credenciales)
	ApiHelpers.RespondJSON(c, 200, credenciales)

	// if err != nil {
	// 	log.Println("Error on list rols")
	// 	ApiHelpers.RespondJSON(c, 404, credenciales)
	// } else {
	// 	log.Println("Success list rols")
	// 	ApiHelpers.RespondJSON(c, 200, credenciales)
	// }
}
