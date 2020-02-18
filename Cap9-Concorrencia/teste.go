package main
	//Concorrência vs Paralelismo
	//Paralelismo = Fazendo 2 coisas exatamente ao mesmo tempo, isso é possível usando 2 processadores
	//Concorrencia = Capacidade de administrar múltiplas tarefas em um único processador
import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}