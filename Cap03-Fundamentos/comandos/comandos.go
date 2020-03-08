package main

import "fmt"

var lang = "GO"

func main() {
	fmt.Printf("Outro %s em %s!!!\n", "programa", lang)
}

//go
//go help
//go env
//Documentação offline: godoc -http=:6060
//Verifica o código: go doc cmd/vet (ex.: go vet comandos.go)
//go build comandos.go
//Executa e compila: go run comandos.go
//Instalar dependencias: go get -u github.com/go-sql-driver/mysql
