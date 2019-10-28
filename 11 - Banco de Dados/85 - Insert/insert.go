package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:trevas@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios (nome) values (?)")
	stmt.Exec("Ayrton")
	stmt.Exec("Cristine")

	res, _ := stmt.Exec("Thais")

	id, _ := res.LastInsertId()
	fmt.Printf("Último ID inserido: %d\n", id)

	linhas, _ := res.RowsAffected()
	fmt.Printf("Quantidade de linhas afetadas: %d\n", linhas)
}
