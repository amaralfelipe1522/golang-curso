package main

import "fmt"

func main() {
	i := 1
	var p *int = nil

	p = &i //Pegando o endereço de memória da variável i e atribuindo ao ponteiro
	*p++   //Asterisco desreferencia, pegando o valor associado ao endereço de memória
	//GO não tem aritméticas de ponteiros. Ex.: p++

	fmt.Println(&i, i, p, *p)
}
