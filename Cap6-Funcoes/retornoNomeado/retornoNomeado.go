package main

import "fmt"

func trocarRetornos(a, b int) (valorB int, valorA int) {
	valorB = b
	valorA = a
	return
}

func main() {
	fmt.Println(trocarRetornos(10, 20))
}
