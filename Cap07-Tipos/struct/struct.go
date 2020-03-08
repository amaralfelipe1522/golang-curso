package main

import "fmt"

type structProduto struct {
	nome     string
	preco    float64
	desconto float64
}

//Método: função com receiver (receptor)
func (p structProduto) produto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	produto1 := structProduto{
		nome:     "Caneca",
		preco:    5.00,
		desconto: 0.05,
	}
	fmt.Println(produto1)
	fmt.Println("Com desconto:", produto1.produto())

	produto2 := structProduto{"Caderno", 20.00, 0.15}
	fmt.Println(produto2)
	fmt.Println("Com desconto:", produto2.produto())
}
