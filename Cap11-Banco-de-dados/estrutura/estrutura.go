package main

//go get -u github.com/go-sql-driver/mysql
//sudo mysql -u root
//ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'Project@1522';
//sudo service mysql restart
//sudo mysql -p
import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //implicitamente necessário, porém não será acessado diretamente
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:Project@1522@/") //Usa indiretamente o drive ao utilizar a string "mysql". Acessando sem ter uma tabela
	if err != nil {
		panic(err)
	}
	defer db.Close() //no fim da função main o banco será encerrado

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
		)`)
}
