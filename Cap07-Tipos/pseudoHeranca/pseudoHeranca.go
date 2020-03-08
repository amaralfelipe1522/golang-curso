package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro //falsa sensação de herança, informando apenas o tipo CARRO é gerado campos anonimos
	turbo bool
}

func main() {
	// Não funciona assim:
	// enzo := ferrari{
	// 	nome:            "Ferrari Enzo",
	// 	velocidadeAtual: 200,
	// 	turbo:           true,
	// }

	enzo := ferrari{}
	enzo.nome = "Ferrari Enzo"
	enzo.velocidadeAtual = 200
	enzo.turbo = true

	fmt.Println(enzo)
}
