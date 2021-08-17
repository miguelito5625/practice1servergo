package Controllers

import (
	"log"
	"practice1servergo/ApiHelpers"
	"practice1servergo/Database/Models"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	var product []Models.Product
	err := Models.GetAllProduct(&product)
	if err != nil {
		log.Println("Error on list products")
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		log.Println("Success list products")
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func AddNewProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.AddNewProduct(&product)
	if err != nil {
		log.Println("Error on insert product:", product)
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		log.Println("Success inserted product:", product)
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func GetOneProduct(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetOneProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func PutOneBook(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetOneProduct(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	}
	c.BindJSON(&product)
	err = Models.PutOneBook(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}

func DeleteBook(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&product, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, product)
	} else {
		ApiHelpers.RespondJSON(c, 200, product)
	}
}
