package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	//converte integer para float64
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)
	//Cuidado, da forma abaixo vai aparecer o elemento correspondente da tabela ASCII
	fmt.Println("Concatenar string com inteiro incorretamente " + string(123))
	fmt.Println("Concatenar string com inteiro corretamente " + strconv.Itoa(123))

	//String para inteiro
	a := "10"
	b := 5
	//essa conversão retorna o valor convertido e um "erro"
	c, _ := strconv.Atoi(a)
	fmt.Println(b + c)

	//String para boolean
	boolEx, _ := strconv.ParseBool("true")
	if boolEx == true {
		fmt.Println(boolEx)
	}
}
