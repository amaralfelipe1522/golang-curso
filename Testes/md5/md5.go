package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func encodeMd5(text string) []byte{
	encode := md5.New()
	io.WriteString(encode,text)
	return encode.Sum(nil)
}

func main() {
	toEncode := "Felipe"
	fmt.Printf("%x",encodeMd5(toEncode))
}