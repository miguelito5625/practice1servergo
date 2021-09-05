package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListCentro_Vacunacion(c *gin.Context) {
	var centro_vacunacion []Models.Centro_Vacunacion
	err := Models.GetAllCentro_Vacunacion(&centro_vacunacion)
	if err != nil {
		log.Println("Error on list rols")
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	} else {
		log.Println("Success list rols")
		ApiHelpers.RespondJSON(c, 200, centro_vacunacion)
	}
}

func AddNewCentro_Vacunacion(c *gin.Context) {
	var centro_vacunacion Models.Centro_Vacunacion
	c.BindJSON(&centro_vacunacion)
	err := Models.AddNewCentro_Vacunacion(&centro_vacunacion)
	if err != nil {
		log.Println("Error on insert centro_vacunacion:", centro_vacunacion)
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	} else {
		log.Println("Success inserted centro_vacunacion:", centro_vacunacion)
		ApiHelpers.RespondJSON(c, 200, centro_vacunacion)
	}
}

func GetOneCentro_Vacunacion(c *gin.Context) {
	id := c.Params.ByName("id")
	var centro_vacunacion Models.Centro_Vacunacion
	err := Models.GetOneCentro_Vacunacion(&centro_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, centro_vacunacion)
	}
}

func PutOneCentro_Vacunacion(c *gin.Context) {
	var centro_vacunacion Models.Centro_Vacunacion
	id := c.Params.ByName("id")
	err := Models.GetOneCentro_Vacunacion(&centro_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	}
	c.BindJSON(&centro_vacunacion)
	err = Models.PutOneCentro_Vacunacion(&centro_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, centro_vacunacion)
	}
}

func DeleteCentro_Vacunacion(c *gin.Context) {
	var centro_vacunacion Models.Centro_Vacunacion
	id := c.Params.ByName("id")
	err := Models.DeleteCentro_Vacunacion(&centro_vacunacion, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, centro_vacunacion)
	} else {
		ApiHelpers.RespondJSON(c, 200, centro_vacunacion)
	}
}
