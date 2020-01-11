package main

import "fmt"

func main() {
	//array possui estrutura homogênea (mesmo tipo) e estática (fixa em qtd de posições)
	var notas [3]float64
	notas[0], notas[1], notas[2] = 7.8, 5.5, 9.5
	fmt.Println(notas)

	total := 0.0

	for i := 0; i < len(notas); i++ {
		total = notas[i]
	}
	media := total / float64(len(notas))

	fmt.Printf("A média é %.2f", media)
}
