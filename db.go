package main

import (
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func createDB() {
	db, _ := sql.Open("sqlite3", "./blog.db")
	defer db.Close()

	stmt, error := db.Prepare("CREATE TABLE IF NOT EXISTS articulos (id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE, title TEXT NOT NULL, content BLOB NOT NULL, date	TEXT NOT NULL)")
	if error != nil {
		fmt.Println("Error creando la tabla [articulos]:", error)
		return
	}
	defer stmt.Close()

	_, error = stmt.Exec()
	if error != nil {
		fmt.Println("Error al ejecutar query para crear la tabla [articulos]:", error)
		return
	}
}