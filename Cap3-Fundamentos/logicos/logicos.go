package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2 //OR exclusivo
	comprarSorvete := trab1 || trab2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv de 50: %t. Tv de 32: %t. Tomou sorvete: %t. Não tomou sorvete: %t.", tv50, tv32, sorvete, !sorvete)
}
