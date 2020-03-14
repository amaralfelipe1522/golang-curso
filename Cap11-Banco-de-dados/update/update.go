package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //implicitamente necessário, porém não será acessado diretamente
)

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Harold", 1)
	stmt.Exec("Danilo", 2)

	stmt, _ = db.Prepare("delete from usuarios where id = ?")
	stmt.Exec(10)
}
