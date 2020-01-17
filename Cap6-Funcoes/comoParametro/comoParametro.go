package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(f func(a, b int) int, a, b int) int {
	return f(a, b)
}

func main() {
	resultado := exec(multiplicacao, 10, 2)
	fmt.Println(resultado)
}
