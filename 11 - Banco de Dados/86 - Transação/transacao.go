package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:trevas@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios (id, nome) values (?, ?)")

	stmt.Exec(4000, "Daniele")
	stmt.Exec(4001, "Vanessa")
	_, err = stmt.Exec(1, "Mateus") // Chave duplicada.
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
