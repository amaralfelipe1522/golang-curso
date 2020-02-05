package main

import "fmt"

func ateZero(num uint) uint {
	switch {
	case num == 0:
		return num
	default:
		fmt.Println("Loop i: ", num)
		ateZeroAnterior := ateZero(num - 1)
		return ateZeroAnterior
	}
}

func main() {
	resultado := ateZero(100)
	fmt.Println("Resultado:", resultado)
}
