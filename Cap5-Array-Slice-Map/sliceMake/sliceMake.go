package main

import "fmt"

func main() {
	s := make([]int, 10) //como o slice não se refere a um array, o compilador cria automaticamente um array
	s[9] = 99
	fmt.Println(s)

	s = make([]int, 10, 20) //cria um slice com 10 posições que se refere a um array interno de 20 posições
	s[9] = 88
	fmt.Println(s, len(s), cap(s)) //Exibe tamanho e capacidade do slice

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(s, len(s), cap(s)) //atualizou automaticamente a capacidade do array interno para 40
}
