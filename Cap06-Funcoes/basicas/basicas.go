package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2) //não printa no console, apenas retorna a string
}

func f5(p1, p2 int) (int, int) {
	soma := p1 + p2
	subt := p1 - p2

	return soma, subt
}

func f6() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Parametro 1", "Parametro 2")
	fmt.Println(f3())
	fmt.Println(f4("Parametro 1", "Parametro 2"))
	fmt.Println(f5(10, 5))
	r1, r2 := f6() //Pode ignorar um dos retornos com _
	fmt.Println(r1, r2)
}
