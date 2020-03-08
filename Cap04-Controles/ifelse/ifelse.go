package main

import "fmt"

func imprimirResultado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com a nota", nota)
	} else {
		fmt.Println("Reprovado com a nota", nota)
	}
}

func main() {
	imprimirResultado(6.4)
}
