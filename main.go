package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/dgrijalva/jwt-go"
)

func main(){
	createDB()
	server := gin.Default()

	server.GET("/", handlerRead)
	server.POST("/", handlerCreate)
	server.PATCH("/:id", handlerUpdate)
	server.DELETE("/:id", handlerDelete)
	
	server.Run()
}