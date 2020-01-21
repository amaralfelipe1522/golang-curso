package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func analisar(a, b int) int {
	result := 0
	if a > b {
		result = subtracao(a, b)
	} else {
		result = soma(a, b)
	}
	return result
}

// func exec(f func(a, b int), a, b int) int {
// 	return f(a, b)
// }

func main() {
	//resultado := exec(analisar, 10, 8)
	resultado := analisar(10, 14)
	fmt.Println(resultado)
}
