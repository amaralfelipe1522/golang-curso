package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"Felipe Amaral": 5000.00,
		"Danilo Reis":   4500.00,
		"Gabriel Fosco": 2600.00,
	}

	funcionarios["Felipe Vitai"] = 3500    //Se já existisse um Felipe Vitai, iria substituir
	delete(funcionarios, "Lucas Natalino") //Não da erro quando o index não existe

	for nome, salario := range funcionarios {
		fmt.Println(nome, salario)
	}
}
