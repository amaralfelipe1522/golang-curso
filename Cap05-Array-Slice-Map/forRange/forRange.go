package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //[...] compilador que define o tamanho do array

	for i, elem := range numeros {
		fmt.Println(i, elem)
	}

	for _, elem := range numeros { // _ ignora um dos parametros
		fmt.Println(elem)
	}

}
