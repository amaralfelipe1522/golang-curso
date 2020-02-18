package main

import "fmt"

func main() {
	//Channel tb é um Tipo da linguagem, assim como float, int, etc
	ch := make(chan int, 1) //2º parametro é um buffer
	//Enviando 10 para o channel
	ch <- 10
	fmt.Println(<-ch)
	ch <- 20
	n := <-ch
	fmt.Println(n)
}