package main

import "fmt"

//Funções variáticas não funcionam com Arrays, apenas com Slices
func imprimir(aprovados ...string) {
	for _, aprovado := range aprovados {
		fmt.Println(aprovado)
	}
}

func main() {
	aprovados := make([]string, 20)
	aprovados[0] = "Felipe"
	aprovados[1] = "Haroldo"
	aprovados[2] = "Gerardo"
	imprimir(aprovados...)
}
