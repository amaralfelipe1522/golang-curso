package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem //sempre que chegar um dado na origem, é encaminhado para o destino
	}
}

// multiplexar = misturar mensagens/dados de canais diferentes em um unico
func multiplexar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

// funcão titulo é da aula anterior (generator)
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
	c := multiplexar(
		titulo("https://www.youtube.com", "https://github.com/amaralfelipe1522"),
		titulo("https://www.google.com", "https://www.uol.com.br"),
	)

	fmt.Println("Primeiros a chegarem nos channels:", <-c, " e ", <-c)
	fmt.Println("Segundos a chegaren:", <-c, " e ", <-c)
}
