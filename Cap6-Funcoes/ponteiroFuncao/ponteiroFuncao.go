package main

import "fmt"

func increment1(n int) {
	n++
	fmt.Println(n)
}

func increment2(n *int) {
	*n++
	fmt.Println(*n)
}

func main() {
	n := 1
	increment1(n)
	fmt.Println(n)

	increment2(&n)
	fmt.Println(n)
}
