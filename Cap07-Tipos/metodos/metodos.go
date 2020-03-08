package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getNome() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setNome(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	pessoa := pessoa{
		nome:      "Felipe",
		sobrenome: "Amaral",
	}
	fmt.Println("Nome completo:", pessoa.getNome())
	pessoa.setNome("Gabriel Fosco")
	fmt.Println("Nome completo:", pessoa.getNome())
}
