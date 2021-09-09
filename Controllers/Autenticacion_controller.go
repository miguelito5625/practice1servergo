package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"
	"practice1servergo/Interfaces"
	"practice1servergo/Services"

	"github.com/gin-gonic/gin"
)

func IniciarSesion(c *gin.Context) {
	var credenciales Interfaces.Credenciales
	c.BindJSON(&credenciales)
	log.Println(credenciales)

	var usuario Models.Usuario
	err := Models.LoginUsuario(&usuario, credenciales.Cui)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, nil, "usuario no existe")
		return
	}

	hashClave := usuario.Clave
	log.Println("LA CLAVE ES: ", hashClave)

	claveCorrecta := Services.CheckPasswordHash(credenciales.Clave, hashClave)

	if !claveCorrecta {
		log.Println("Error autenticacion clave incorrecta")
		ApiHelpers.RespondJSON(c, 404, nil, "Error: clave incorrecta")
		return
	}
	token, _ := Services.CreateToken(credenciales)
	ApiHelpers.RespondJSON(c, 200, token, "ok")

}
