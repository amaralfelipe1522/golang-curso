package main

import "fmt"

func (pkm Pokemon) atacar(a int) {
	fmt.Println("Charmander usou " + pkm.moveSet[a])
}

func (pkm Pokemon) sofreuDano(dano int) int {
	pkm.hp = pkm.hp - dano
	if pkm.hp < 1 {
		pkm.hp = 0
		fmt.Println("Charmander recebeu", dano, "de dano. Charmander desmaiou...")
	} else {
		fmt.Println("Charmander recebeu", dano, "de dano. Ainda possui", pkm.hp, "de HP.")
	}
	return pkm.hp
}

// func (pkm Pokemon, ptn Potion) recuperarHP(potionType)
