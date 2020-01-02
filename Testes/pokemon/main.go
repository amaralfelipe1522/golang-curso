package main

import (
	"fmt"
)

// type Pokemon struct {
// 	name    string
// 	level   int
// 	hp      int
// 	moveSet [4]string
// }

// func (pkm Pokemon) atacar(a int) {
// 	fmt.Println("Charmander usou " + pkm.moveSet[a])
// }

// func (pkm Pokemon) sofreuDano(dano int) int {
// 	pkm.hp = pkm.hp - dano
// 	fmt.Println("Charmander recebeu", dano, "de dano. Ainda possui", pkm.hp, "de HP.")
// 	return pkm.hp
// }

func main() {

	pokemon01 := Pokemon{
		name:    "Charmander",
		level:   12,
		hp:      200,
		moveSet: [4]string{"Tackle", "Quick Attack", "Ember", "Fire Spin"},
	}

	// potions := Potion{
	// 	potion:      20,
	// 	superPotion: 50,
	// 	megaPotion:  100,
	// 	hyperPotion: 200,
	// }

	fmt.Println(pokemon01)
	// fmt.Println(potions)
	pokemon01.atacar(2)
	pokemon01.hp = pokemon01.sofreuDano(10)
	fmt.Println(pokemon01)
	pokemon01.hp = pokemon01.sofreuDano(190)
	fmt.Println(pokemon01)
}
