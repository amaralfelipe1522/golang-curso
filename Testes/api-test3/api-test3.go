package main

import "fmt"

func setURL(pokemon string) string {
	url := "https://pokeapi.co/api/v2/pokemon/"
	return url + pokemon + "/"
}

func main() {
	url := setURL("pikachu")
	fmt.Println(url)
}
