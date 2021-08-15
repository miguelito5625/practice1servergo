package User_routes

import (
	"github.com/gin-gonic/gin"
)

func getAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user login",
	})
}

func create(c *gin.Context) {

	// var login structs.USER_STRUCT
	// c.BindJSON(&login)

	// fmt.Println("DATOS :")
	// fmt.Println(login)
	c.JSON(200, gin.H{
		// "user":    login,
		"message": "user created",
		"status":  "ok",
	})
}

func Routes(route *gin.Engine) {
	userRoutes := route.Group("/user")
	userRoutes.GET("", getAll)
	userRoutes.POST("", create)
}
