// Google I/O 2012 - Go Concurrency Patterns
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// lista de urls como parametro que retorna um channel somente leitura
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>") // cria a regra de regex
			c <- r.FindStringSubmatch(string(html))[1]       // converte o conteúdo html para string e aplica o regex
		}(url) // dispara a execução da função anonima passando url como parametro
	}
	return c
}

func main() {
	t1 := titulo("https://www.youtube.com", "https://github.com/amaralfelipe1522")
	t2 := titulo("https://www.google.com", "https://www.uol.com.br")

	fmt.Println("Primeiros a chegarem nos channels:", <-t1, " e ", <-t2)
	fmt.Println("Segundos a chegaren:", <-t1, " e ", <-t2)
}
