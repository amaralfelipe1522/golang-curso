package main

import (
	"fmt"
	"time"
)

func identTipo(dado interface{}) string { //recebe argumento genérico: interface{}
	switch dado.(type) {
	case int:
		return "É um valor inteiro"
	case float64, float32:
		return "É um float"
	case bool:
		return "É um boolean"
	case string:
		return "É uma string"
	case func():
		return "É uma função"
	default:
		return "Tipo desconhecido"
	}
}

func main() {
	fmt.Println(identTipo(2))
	fmt.Println(identTipo(2.1))
	fmt.Println(identTipo(true))
	fmt.Println(identTipo("Teste"))
	fmt.Println(identTipo(func() {}))
	fmt.Println(identTipo('a'))
	fmt.Println(identTipo(time.Now()))
}
