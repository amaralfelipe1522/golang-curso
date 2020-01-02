package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String: ", "Banana" == "Banana")
	fmt.Println("!= :", 3 != 2)
	fmt.Println("< :", 2 < 3)
	fmt.Println("> :", 2 > 3)
	fmt.Println("<= :", 3 <= 3)
	fmt.Println("=> :", 2 >= 3)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	//não leva em consideração a alocação da memória, mas sim os valores
	//quando se usa Ponteiros a história é outra
	fmt.Println(d1 == d2)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		name string
		age  int
	}

	pessoa1 := Pessoa{"Felipe", 12}
	pessoa2 := Pessoa{"Felipe", 12}

	fmt.Println("Pessoas: ", pessoa1 == pessoa2)
}
