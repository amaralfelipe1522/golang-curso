package main

import "fmt"

func main() {
	aprovados := make(map[int]string) //chave int, valor string

	aprovados[40174303840] = "Felipe"
	aprovados[23515222349] = "Gabriel"
	fmt.Println(aprovados)
	aprovados[23515222349] = "Gabs"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println("CPF:", cpf, "\nNome:", nome)
	}

	delete(aprovados, 23515222349)
	fmt.Println(aprovados)
}
