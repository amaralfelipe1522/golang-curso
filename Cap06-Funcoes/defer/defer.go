package main

import "fmt"

//indicar a palavra reservada DEFER faz com que atrase a execução de um método até um momento antes do retorno da função ocorrer

func deferExample(v1, v2 int) int {
	defer fmt.Println("Fim")
	fmt.Println("Inicio")
	resultado := 0
	if v1 > v2 {
		resultado = v1 - v2
	} else {
		resultado = v2 - v1
	}
	return resultado
}

func main() {
	fmt.Println(deferExample(3, 4))
}
