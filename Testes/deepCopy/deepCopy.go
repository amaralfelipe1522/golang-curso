package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := map[string]map[string]map[string]interface{}{
		"Felipe": {
			"Dados Pessoais": {
				"CPF":    40174303840,
				"dtNasc": "25/09/1991",
			},
			"Experiência Profissional": {
				"Magna": "Até o momento",
				"Ultra": "03/09/2019",
			},
		},
	}

	fmt.Println("Map original:", map1)

	dtMap, _ := json.Marshal(map1) //omitindo o erro, se houver
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	fmt.Println(dtMap)

	jsonToString := string(dtMap)
	fmt.Println("Convertido em JSON:", jsonToString)

	map2 := make(map[string]interface{})

	json.Unmarshal([]byte(jsonToString), &map2) //tb pode dar erro, mas eu não tratei
	fmt.Println("Novo map:", map2)
}
