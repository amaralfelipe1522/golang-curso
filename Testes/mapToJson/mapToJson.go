package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := map[string]int{
		"Felipe":  123,
		"Gabriel": 456,
	}

	fmt.Println(map1)

	dtMap, _ := json.Marshal(map1) //omitindo o erro, se houver
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(dtMap)

	jsonStr := string(dtMap)
	fmt.Println(jsonStr)
}
