package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.129212
	fmt.Println("Valor de X é", x, "!!")

	//Converte para String para ser concatendo no PRINT
	xs := fmt.Sprint(x)
	fmt.Println("Valor de X é " + xs)
	//%f = float
	//%s = string
	//%d = integer
	//%t = boolean
	//%.2f = Formata para 2 casas decimais
	//%v = imprime varios tipos
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
}
