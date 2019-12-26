package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	type Character struct {
		name string
		age  int
	}

	red := Character{
		name: "Red",
		age:  12,
	}

	blue := Character{
		name: "Blue",
		age:  13,
	}

	fmt.Println(red, blue)
}
