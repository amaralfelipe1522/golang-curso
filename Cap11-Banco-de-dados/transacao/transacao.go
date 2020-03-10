package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //implicitamente necessário, porém não será acessado diretamente
)

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/cursogo") //database criada na aula estrutura.go
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values(?,?)")

	stmt.Exec(10, "Danilo")
	stmt.Exec(11, "Gerardo")
	_, err = stmt.Exec(12, "Celta")

	if err != nil {
		tx.Rollback() //Se houver algum erro, aplica-se o ROLLBACK
		log.Fatal(err)
	}

	tx.Commit()
}
