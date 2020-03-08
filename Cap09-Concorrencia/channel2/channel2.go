package main

import (
	"fmt"
	"time"
)

func multiplica(num int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * num //Envia para o canal

	time.Sleep(time.Second)
	ch <- 3 * num

	time.Sleep(time.Second * 3)
	ch <- 4 * num
}

func main() {
	ch := make(chan int)
	go multiplica(2, ch)
	fmt.Println("Iniciando..")

	a, b := <-ch, <-ch
	fmt.Println("A e B recebidos..")
	fmt.Println(a,b)

	fmt.Println("Falta um..")
	fmt.Println(<-ch)
	fmt.Println("Último recebido!")

	fmt.Println(<-ch) //deadlock
}