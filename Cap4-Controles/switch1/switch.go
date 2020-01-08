package main

import "fmt"

//Switch não é verdadeiro e falso

//Diferente do Switch de outras linguagens, em GO não precisa de um BREAK
//Uma alternativa para isso é usar FALLTHROUGH
func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "Nota é A"
	case 8, 7:
		return "Nota é B"
	case 6, 5:
		return "Nota é C"
	case 4, 3:
		return "Nota é D"
	case 2, 1, 0:
		return "Nota é E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(1))
}
