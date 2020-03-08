package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	fmt.Println(r)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //Bloco de inicialização que tb pode ser usado no Switch
		fmt.Println(i, "é maior que cinco")
	} else {
		fmt.Println(i, "é menor que cinco") //variável só pode ser acessada dentro do IF/ELSE
	}
}
