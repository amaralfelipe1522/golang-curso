package main

import (
	"fmt"
	"math"
	"reflect" //identifica o tipo do dado
)

func main() {
	//integer
	fmt.Println(1, 2, 1000)
	fmt.Println("Verifica o tipo do valor:", reflect.TypeOf(32000))

	//sem sinal (só positivos): uint8, uint16, uint32 e uint64
	//byte = uint8
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	//com sinal: int8, int16, int32, int64
	maxInt := math.MaxInt64
	fmt.Println(maxInt)
	fmt.Println(reflect.TypeOf(maxInt))

	var varUnicode rune = 'a' //mostra qual o valor referente a tabela unicode (int32)
	fmt.Println("unicode", varUnicode)
	fmt.Println(reflect.TypeOf(varUnicode))

	//float32 e float64
	var varFloat float32 = 4.99
	fmt.Println(reflect.TypeOf(varFloat))
	fmt.Println(reflect.TypeOf(4.99)) //float64

	//boolean
	boo := true
	fmt.Println(reflect.TypeOf(boo))
	fmt.Println(!boo)

	//string
	varString := "Hello world"
	fmt.Println(varString + "!!")
	fmt.Println(len(varString))

	//string com múltiplas linhas
	multiString := `Hello
	World`
	fmt.Println(multiString)

	//char??
	//aceita aspas simples, mas no fim das contas é um int32
	varChar := 'a'
	fmt.Println(varChar)
	fmt.Println(reflect.TypeOf(varChar))
}
