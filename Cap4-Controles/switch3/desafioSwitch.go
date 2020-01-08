package main

// Desafio para refatorar o código do arquivo ifelseif.go
import "fmt"

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 6 && n < 8:
		return "C"
	case n >= 4 && n < 6:
		return "D"
	case n >= 0 && n < 4:
		return "E"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Println("A nota é:", notaParaConceito(10))
}
