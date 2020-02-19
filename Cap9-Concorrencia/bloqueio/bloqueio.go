package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	time.Sleep(time.Second)
	ch <- 1 // bloqueante
	fmt.Println("Leu")
}

func main() {
	ch := make(chan int)
	go rotina(ch)
	fmt.Println("Vai ler")
	//fmt.Println(<-ch)
	
	fmt.Println(<-ch) // bloqueante
	fmt.Println("Fim")
}