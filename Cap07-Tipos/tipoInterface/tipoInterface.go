package main

import "fmt"

func main() {
	var coisa interface{}
	coisa = 3
	fmt.Println(coisa)
	coisa = true
	fmt.Println(coisa)
	coisa = "Tipagem dinamica"
	fmt.Println(coisa)
}
