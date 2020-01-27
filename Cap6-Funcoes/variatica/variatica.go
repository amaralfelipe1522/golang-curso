package main

import "fmt"

func media(notas ...float64) float64 {
	total := 0.0
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

func main() {
	fmt.Printf("A média é %.2f", media(5.5, 6.5, 5.5))
}
