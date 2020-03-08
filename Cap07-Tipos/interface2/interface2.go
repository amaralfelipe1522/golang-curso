package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	//manipulando a nível do Tipo
	enzo := ferrari{"Ferrari Enzo", false, 0}
	fmt.Println(enzo)
	enzo.ligarTurbo()
	fmt.Println(enzo)

	//ou manipulando a nível de interface
	var maranello esportivo = &ferrari{"Ferrari Maranello", false, 0}
	fmt.Println(maranello)
	maranello.ligarTurbo()
	fmt.Println(maranello)
}
