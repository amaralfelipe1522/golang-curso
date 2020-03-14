package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //implicitamente necessário, porém não será acessado diretamente
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios where id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
