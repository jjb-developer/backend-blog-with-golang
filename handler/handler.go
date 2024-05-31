package handler

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"../model"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(){
	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS articulos (
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		title TEXT NOT NULL,
		content BLOB NOT NULL,
		date TEXT NOT NULL
		)`)
	stmt.Exec()
	stmt.Close()
}

func Read(ctx *gin.Context){
	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	rows, error := db.Query(`SELECT * FROM articulos`)
	if error != nil {
		fmt.Println("Ha ocurrido un error:", error)
		return
	}

	var articulos []model.Articulo
	for rows.Next() {
		var articulo model.Articulo
		rows.Scan(&articulo.Id, &articulo.Title, &articulo.Content, &articulo.Date)
		articulos = append(articulos, articulo)
	}

	ctx.JSON(200, gin.H{"result": articulos})
}


func Create(ctx *gin.Context){
	var articulo model.BodyPostRequest
	ctx.BindJSON(&articulo)
	articulo.Date = fmt.Sprintf("%d", time.Now().Unix())

	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	ok, error := db.Exec(`INSERT INTO articulos (title,content,date) VALUES (?,?,?)`, articulo.Title, articulo.Content, articulo.Date)
	if error != nil {
		fmt.Println("Ha ocurrido un error:", error)
		return
	}
	fmt.Println("articulo creado en DB exitosamente:", ok)
	ctx.JSON(200, gin.H{"message": "articulo creado en DB exitosamente!."})
}


func UpdateOne(ctx *gin.Context){
	ctx.JSON(200, gin.H{"message": "ok!."})
}


func Update(ctx *gin.Context){
	ctx.JSON(200, gin.H{"message": "ok!."})
}


func Delete(ctx *gin.Context){
	ctx.JSON(200, gin.H{"message": "ok!."})
}