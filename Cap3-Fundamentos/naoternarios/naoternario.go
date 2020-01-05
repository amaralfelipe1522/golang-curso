package main

import "fmt"

//Não existem operadores ternários em GO
//Alternativa abaixo:

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
	//return nota >= 6 "Aprovado" : "Reprovado"
}

func main() {
	resultado := obterResultado(5.9)
	fmt.Println(resultado)
}
