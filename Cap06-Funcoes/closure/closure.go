package main

//função respeita a origem
import "fmt"

func closureEx() func() {
	x := 10
	printX := func() {
		fmt.Println(x)
	}
	return printX
}

func main() {
	x := 20
	fmt.Println(x)

	callFunc := closureEx()
	callFunc()
}
