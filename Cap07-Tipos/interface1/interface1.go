package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("Produto: %s custa R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var pessoa1 imprimivel = pessoa{"Felipe", "Amaral"}
	var produto1 imprimivel = produto{"Caneca", 5.00}

	fmt.Println(pessoa1.toString())
	fmt.Println(produto1.toString())
	//ou
	imprimir(pessoa1)
	imprimir(produto1)
	//ou
	imprimir(pessoa{"Gabriel", "Fosco"})
	imprimir(produto{"Caderno", 10.00})
}
