package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//switch sem valor associado vai buscar um valor TRUE
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia.")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
