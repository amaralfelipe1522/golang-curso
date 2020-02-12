package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	//Convenção: identificador começando com letra maiúscula significa ter uma semantica Pública
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct to json
	produto1 := produto{1, "Notebook Dell", 2000.00, []string{"Notebook", "Eletronico"}}
	fmt.Println(produto1)
	prodJSON, _ := json.Marshal(produto1)
	fmt.Println(string(prodJSON))

	//json to struct
	var produto2 produto
	prodStruct := `{"id":2,"nome":"Notebook LG","preco":1500,"tags":["Notebook","Eletronico"]}`
	json.Unmarshal([]byte(prodStruct), &produto2)
	fmt.Println(produto2)
}
