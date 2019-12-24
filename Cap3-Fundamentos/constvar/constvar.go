package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2
	fmt.Println(PI)
	fmt.Println(raio)

	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferênce ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	var e, f bool = true, false

	g, h, i := true, 12, "Epa"

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
