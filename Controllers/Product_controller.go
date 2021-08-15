package Controllers

import (
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Product
	err := Models.GetAllBook(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func AddNewProduct(c *gin.Context) {
	var book Models.Product
	c.BindJSON(&book)
	err := Models.AddNewProduct(&book)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Product
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Product
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, book)
	} else {
		ApiHelpers.RespondJSON(c, 200, book)
	}
}
