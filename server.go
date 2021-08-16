package main

import (
	"io"
	"log"
	"os"
	"practice1servergo/Database"
	"practice1servergo/Database/Migrations"
	"practice1servergo/Routes/Product_routes"
	"practice1servergo/Routes/User_routes"

	"github.com/gin-gonic/gin"
)

func main() {

	Database.ConnectDataBase()
	Migrations.MigrateAll()

	logfile, _ := os.Create("server.log")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("Welcome to API GO")
		c.JSON(200, gin.H{
			"message": "Hello!!!",
		})
	})
	User_routes.Routes(r)
	Product_routes.Routes(r)
	r.Run(":3000")
}
