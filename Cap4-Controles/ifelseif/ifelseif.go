package main

import "fmt"

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 6 && n < 8 {
		return "C"
	} else if n < 6 && n >= 0 {
		return "D"
	} else {
		return "Nota inválida"
	}
}

func main() {
	resultado := notaParaConceito(-1)
	fmt.Println("A nota é:", resultado)
}
