package main

import (
	"fmt"
)

func atacar(pkm string) {
	fmt.Println("Charmander usou " + pkm)
}

func sofreuDano(hp, dano int) int {
	resultado := hp - dano
	fmt.Println("Charmander recebeu", dano, "de dano. Ainda possui", resultado, "de HP.")
	return resultado
}

func main() {
	type Pokemon struct {
		name    string
		level   int
		hp      int
		moveSet [4]string
	}

	pokemon01 := Pokemon{
		name:    "Charmander",
		level:   12,
		hp:      200,
		moveSet: [4]string{"Tackle", "Quick Attack", "Ember", "Fire Spin"},
	}

	//fmt.Println(pokemon01)

	atacar(pokemon01.moveSet[3])
	sofreuDano(pokemon01.hp, 10)
}
