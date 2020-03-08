package main

import "fmt"

func buffando(ch chan int, valor int) {
	for i:=0; i<valor; i++ {
		ch <- i+1
	}
}

func main() {
	valor := 5
	ch := make(chan int, valor)

	go buffando(ch,valor)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch) //Deadlock pois atingiu o limite de buffer do channel
}