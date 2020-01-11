package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice não possui tamanho fixo

	fmt.Println(a1, s1)
	//slice não é um array
	//slice define um pedaço de um array
	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] //Slice apontando para os elementos do indice 1 até 2 (assim como 1 até <3)
	fmt.Println(s2)

	s3 := a2[:2] //aponta para o inicio do array
	fmt.Println(s3)

	s4 := s2[:1] //slice apontando para um slice
	fmt.Println(s4)
}
