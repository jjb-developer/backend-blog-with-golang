package main

import (
	"./handler"
	"github.com/gin-gonic/gin"

	//"time"
	//"github.com/dgrijalva/jwt-go"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main(){
	//handler.CreateDB()
	server := gin.Default()
	//server.GET("/createDB", handler.CreateDB)
	server.GET("/", handler.Read)
	server.POST("/", handler.Create)
	//server.PATCH("/:id", handler.UpdateOne)
	//server.PUT("/", handler.Update)
	//server.DELETE("/", handler.Delete)
	server.Run(":8080")
}