package main

import (
	"fmt"

	"time"
	"github.com/gin-gonic/gin"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func handlerRead(ctx *gin.Context){
	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	rows, error := db.Query("SELECT * FROM articulos")
	if error != nil {
		fmt.Println("Error procesando una Query de selección:", error)
		return
	}
	defer rows.Close()

	var articulos []Article
	for rows.Next() {
		var articulo Article
		rows.Scan(&articulo.Id, &articulo.Title, &articulo.Content, &articulo.Date)
		articulos = append(articulos, articulo)
	}
	ctx.JSON(200, gin.H{
		"result": articulos,
	})
}


func handlerCreate(ctx *gin.Context){
	var articulo BodyRequest
	ctx.BindJSON(&articulo)
	articulo.Date = fmt.Sprintf("%d", time.Now().Unix())

	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	stmt, error := db.Prepare("INSERT INTO articulos (title, content, date) VALUES (?,?,?)")
	if error != nil {
		fmt.Println("Error preparando la info del articulo para agregarlo en la tabla [articulos]:", error)
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec(articulo.Title, articulo.Content, articulo.Date)
	if error != nil {
		fmt.Println("Error al ejecutar Query de escritura de inserción de datos en la DB:", error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Articulo creado exitosamente!.",
	})
}


func handlerUpdate(ctx *gin.Context){
	var articulo BodyRequest
	ctx.BindJSON(&articulo)
	articulo.Date = fmt.Sprintf("%d", time.Now().Unix())

	id := ctx.Param("id")
	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	stmt, error := db.Prepare("UPDATE articulos SET title=?, content=?, date=? WHERE id=?")
	if error != nil {
		fmt.Println("Error preparando la info para actualizar items dentro de la tabla [articulos]:", error)
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec(articulo.Title, articulo.Content, articulo.Date, id)
	if error != nil {
		fmt.Println("Error al ejecutar Query de escritura de actualizacion de datos en la DB::", error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Articulo actualizado exitosamente!.",
	})
}


func handlerDelete(ctx *gin.Context){
	id := ctx.Param("id")

	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	stmt, error := db.Prepare("DELETE FROM articulos WHERE id=?")
	if error != nil {
		fmt.Println("Error preparando la info para eliminar items dentro de la tabla [articulos]:", error)
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec(id)
	if error != nil {
		fmt.Println("Error al ejecutar Query de escritura en la eliminacion de datos en la DB::", error)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Articulo eliminado exitosamente!.",
	})
}