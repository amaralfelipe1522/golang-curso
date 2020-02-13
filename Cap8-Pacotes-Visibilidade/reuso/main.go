package main

import (
	"fmt"

	"github.com/amaralfelipe1522/area"
	/*
		package area

		import "math"

		//Pi é uma proporção númerica definida pela relação entre o perímetro de uma circunferência e seu diâmetro
		const Pi = 3.1416

		//Circ é responsável por calcular a área
		func Circ(raio float64) float64 {
			return math.Pow(raio, 2) * Pi
		}

		//Rect é responsável por calcular a área de um retangulo
		func Rect(base, altura float64) float64 {
			return base * altura
		}

		//_TrianguloEq é privada, não será visível em outro pacote
		func _TrianguloEq(base, altura float64) float64 {
			return (base * altura) / 2
		}
	*/)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	//_TrianguloEq não é acessado
}
