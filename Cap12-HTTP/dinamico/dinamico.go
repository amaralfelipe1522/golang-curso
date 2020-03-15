package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	// Formato de data:
	// 01: mês / 02: dia / 03: hora / 04: minuto / 05: segundo / 2006: 4 dígitos para o ano
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", s) // Fprintf printa a mensagem no ResponseWriter
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
