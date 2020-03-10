package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //implicitamente necessário, porém não será acessado diretamente
)

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/cursogo") //database criada na aula estrutura.go
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)") //O valor será inserido no lugar da interrogação
	stmt.Exec("Felipe")
	stmt.Exec("João")

	res, _ := stmt.Exec("Gabriel")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
