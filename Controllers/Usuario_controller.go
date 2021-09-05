package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"
	"practice1servergo/Services"

	"github.com/gin-gonic/gin"
)

func ListUsuario(c *gin.Context) {
	var usuario []Models.Usuario
	err := Models.GetAllUsuario(&usuario)
	if err != nil {
		log.Println("Error on list rols")
		ApiHelpers.RespondJSON(c, 404, usuario)
	} else {
		log.Println("Success list rols")
		ApiHelpers.RespondJSON(c, 200, usuario)
	}
}

func AddNewUsuario(c *gin.Context) {
	var usuario Models.Usuario
	c.BindJSON(&usuario)
	usuario.Clave, _ = Services.HashPassword(usuario.Clave)
	err := Models.AddNewUsuario(&usuario)
	if err != nil {
		log.Println("Error on insert usuario:", usuario)
		ApiHelpers.RespondJSON(c, 404, usuario)
	} else {
		log.Println("Success inserted usuario:", usuario)
		ApiHelpers.RespondJSON(c, 200, usuario)
	}
}

func GetOneUsuario(c *gin.Context) {
	id := c.Params.ByName("id")
	var usuario Models.Usuario
	err := Models.GetOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario)
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario)
	}
}

func PutOneUsuario(c *gin.Context) {
	var usuario Models.Usuario
	id := c.Params.ByName("id")
	err := Models.GetOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario)
	}
	c.BindJSON(&usuario)
	err = Models.PutOneUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario)
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario)
	}
}

func DeleteUsuario(c *gin.Context) {
	var usuario Models.Usuario
	id := c.Params.ByName("id")
	err := Models.DeleteUsuario(&usuario, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, usuario)
	} else {
		ApiHelpers.RespondJSON(c, 200, usuario)
	}
}
