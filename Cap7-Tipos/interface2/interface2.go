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
	enzo := ferrari{"Ferrari Enzo", false, 0}
	fmt.Println(enzo)
	enzo.ligarTurbo()
	fmt.Println(enzo)

	//ou
	var maranello esportivo = &ferrari{"Ferrari Maranello", false, 0}
	fmt.Println(maranello)
	maranello.ligarTurbo()
	fmt.Println(maranello)
}
