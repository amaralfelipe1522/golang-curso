package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode ter mais métodos
}

type bmw struct {
	nome string
}

func (b bmw) ligarTurbo() {
	fmt.Println(b.nome, "ligou o turbo.")
}

func (b bmw) fazerBaliza() {
	fmt.Println(b.nome, "fez baliza.")
}

func main() {
	var bwm1 esportivoLuxuoso = &bmw{"BMW 1"}
	bwm1.fazerBaliza()
	bwm1.ligarTurbo()
}
