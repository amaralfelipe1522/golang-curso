package main

import "fmt"

func ateZero(num int) (int, error) {
	switch {
	case num < 0:
		return num, fmt.Errorf("Erro: valor %d é negativo", num)
	case num == 0:
		return num, nil
	default:
		fmt.Println("Loop i: ", num)
		ateZeroAnterior, _ := ateZero(num - 1)
		return ateZeroAnterior, fmt.Errorf("Loop f")
	}
}

func main() {
	resultado, err := ateZero(100)
	if err != nil {
		fmt.Println("Resultado:", resultado, ". Erro:", err)
	} else {
		fmt.Println("Resultado:", resultado)
	}
}
