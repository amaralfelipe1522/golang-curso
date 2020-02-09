package main

import "fmt"

type item struct {
	itemID int
	preco  float64
	qtd    int
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) carrinho() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 111,
		itens: []item{
			item{
				itemID: 1,
				preco:  10.00,
				qtd:    10,
			},
			item{
				itemID: 2,
				preco:  20.00,
				qtd:    20,
			},
			item{
				itemID: 3,
				preco:  30.00,
				qtd:    30,
			},
		},
	}
	fmt.Printf("Valor total do carrinho é R$%.2f", pedido.carrinho())
}
