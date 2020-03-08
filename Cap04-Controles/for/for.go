package main

import (
	"fmt"
	"time"
)

func main() {
	//Ex 01 - Tipo WHILE
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//Ex 02 - FOR normal
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d", i)
	}
	//Ex 03 - Misturando estruturas de controle
	fmt.Println("Misturando..")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é impar")
		}
	}
	//Laço infinito
	for {
		fmt.Println("Para sempre..")
		time.Sleep(time.Second * 2) //imprimi a cada 2 segundos
	}
}
