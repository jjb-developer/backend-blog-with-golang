package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/dgrijalva/jwt-go"

	//"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

func main(){
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})
	
	server.Run()
}