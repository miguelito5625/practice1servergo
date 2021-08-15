package main

import (
	"practice1servergo/Database"
	"practice1servergo/Database/Migrations"
	"practice1servergo/Routes/Product_routes"
	"practice1servergo/Routes/User_routes"

	"github.com/gin-gonic/gin"
)

func main() {

	Database.ConnectDataBase()
	Migrations.MigrateAll()

	// file, fileErr := os.Create("server.log")
	// if fileErr != nil {
	// 	fmt.Println(fileErr)
	// 	return
	// }
	// gin.DefaultWriter = file

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!!!",
		})
	})
	User_routes.Routes(r)
	Product_routes.Routes(r)
	r.Run(":3000")
}
