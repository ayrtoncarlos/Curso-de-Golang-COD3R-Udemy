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

	// UPDATE
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Ayrton Andrade", 1)
	stmt.Exec("Cristine Diniz", 2)
	stmt.Exec("Thais Dantas", 3)

	// DELETE
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")

	stmt2.Exec(2000)

}
