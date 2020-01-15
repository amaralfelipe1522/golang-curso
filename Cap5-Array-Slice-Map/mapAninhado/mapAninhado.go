package main

import "fmt"

func main() {
	m1 := map[int]map[string]map[string]int{
		1: {
			"Felipe": {
				"CPF":  40174303840,
				"Nasc": 25091991,
			},
		},
		2: {
			"Danilo": {
				"CPF":  39835478940,
				"Nasc": 16071990,
			},
			"Vitai": {
				"CPF":  45455282278,
				"Nasc": 16091994,
			},
		},
	}

	for id, nome := range m1 {
		fmt.Println(id, nome)
		for cpf, nasc := range nome {
			fmt.Println(cpf, nasc)
		}
	}

	// delete(m1[2], "Vitai") //Deleta um sub valor

	// for id, nome := range m1 {
	// 	fmt.Println(id, nome)
	// 	for cpf, nasc := range nome {
	// 		fmt.Println(cpf, nasc)
	// 	}
	// }
}
