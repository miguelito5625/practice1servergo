package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListVacuna(c *gin.Context) {
	var vacuna []Models.Vacuna
	err := Models.GetAllVacuna(&vacuna)
	if err != nil {
		log.Println("Error on list rols")
		ApiHelpers.RespondJSON(c, 404, vacuna)
	} else {
		log.Println("Success list rols")
		ApiHelpers.RespondJSON(c, 200, vacuna)
	}
}

func AddNewVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	c.BindJSON(&vacuna)
	err := Models.AddNewVacuna(&vacuna)
	if err != nil {
		log.Println("Error on insert vacuna:", vacuna)
		ApiHelpers.RespondJSON(c, 404, vacuna)
	} else {
		log.Println("Success inserted vacuna:", vacuna)
		ApiHelpers.RespondJSON(c, 200, vacuna)
	}
}

func GetOneVacuna(c *gin.Context) {
	id := c.Params.ByName("id")
	var vacuna Models.Vacuna
	err := Models.GetOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna)
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna)
	}
}

func PutOneVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	id := c.Params.ByName("id")
	err := Models.GetOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna)
	}
	c.BindJSON(&vacuna)
	err = Models.PutOneVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna)
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna)
	}
}

func DeleteVacuna(c *gin.Context) {
	var vacuna Models.Vacuna
	id := c.Params.ByName("id")
	err := Models.DeleteVacuna(&vacuna, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vacuna)
	} else {
		ApiHelpers.RespondJSON(c, 200, vacuna)
	}
}
