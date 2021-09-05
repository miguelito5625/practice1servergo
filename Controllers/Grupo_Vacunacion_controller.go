package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListGrupo_Vacunacion(c *gin.Context) {
	var grupo_vacunacion []Models.Grupo_Vacunacion
	err := Models.GetAllGrupo_Vacunacion(&grupo_vacunacion)
	if err != nil {
		log.Println("Error on list rols")
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	} else {
		log.Println("Success list rols")
		ApiHelpers.RespondJSON(c, 200, grupo_vacunacion)
	}
}

func AddNewGrupo_Vacunacion(c *gin.Context) {
	var grupo_vacunacion Models.Grupo_Vacunacion
	c.BindJSON(&grupo_vacunacion)
	err := Models.AddNewGrupo_Vacunacion(&grupo_vacunacion)
	if err != nil {
		log.Println("Error on insert grupo_vacunacion:", grupo_vacunacion)
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	} else {
		log.Println("Success inserted grupo_vacunacion:", grupo_vacunacion)
		ApiHelpers.RespondJSON(c, 200, grupo_vacunacion)
	}
}

func GetOneGrupo_Vacunacion(c *gin.Context) {
	id := c.Params.ByName("id")
	var grupo_vacunacion Models.Grupo_Vacunacion
	err := Models.GetOneGrupo_Vacunacion(&grupo_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, grupo_vacunacion)
	}
}

func PutOneGrupo_Vacunacion(c *gin.Context) {
	var grupo_vacunacion Models.Grupo_Vacunacion
	id := c.Params.ByName("id")
	err := Models.GetOneGrupo_Vacunacion(&grupo_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	}
	c.BindJSON(&grupo_vacunacion)
	err = Models.PutOneGrupo_Vacunacion(&grupo_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, grupo_vacunacion)
	}
}

func DeleteGrupo_Vacunacion(c *gin.Context) {
	var grupo_vacunacion Models.Grupo_Vacunacion
	id := c.Params.ByName("id")
	err := Models.DeleteGrupo_Vacunacion(&grupo_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, grupo_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, grupo_vacunacion)
	}
}
