package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public")) // lê a pasta public a partir do diretório atual
	http.Handle("/", fs)                      // Quando chegar uma requisição na raiz da minha aplicação, automaticamente passe para o handle (fs)

	log.Println("Executando na porta 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
