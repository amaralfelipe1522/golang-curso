package main

import (
    b64 "encoding/base64"
    "fmt"
)

func encode64(text string) string {
	textEncode := b64.StdEncoding.EncodeToString([]byte(text))
	return textEncode
}

func decode64(text string) string {
	textDecode, _ := b64.StdEncoding.DecodeString(text)
	return string(textDecode) //converte bytes para string
}

func main() {
	example := "Felipe"
	encodado := encode64(example)
	fmt.Println(encodado)

	decodado := decode64(encodado)
	fmt.Println(decodado)
}