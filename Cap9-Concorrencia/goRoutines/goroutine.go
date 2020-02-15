package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteranção de número %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//Executa de forma sequencial (serial), sem concorrência:
	// fale("João", "Ô diabo", 4)
	// fale("João2", "HEHEHE", 4)

	//Criando uma Go Routine
	go fale("João", "Ô diabo", 5)
	go fale("João2", "HEHEHE", 5)
	time.Sleep(time.Second * 3) //Go Routine terá 3 segundos para executar antes da função Main ser encerrada
}
