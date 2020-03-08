package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randObject() string {
	objetos := [138]string{"Agulha", "Alfinete", "Anel", "Anzol", "Andador", "Apontador", "Almofada", "Abajour", "Bola", "Balão", "Botas", "Bexiga", "Borracha Escolar", "CD", "Colher", "Caneta", "Camisa", "Camiseta", "Calça", "Castiçal", "Copo", "Carimbo", "Dado", "Dedal", "Dicionário", "Diário", "DVD", "Dentadura", "Escada", "Escova", "Espelho", "Espada", "Esmalte", "Faca", "Facão", "Folha", "Furadeira", "Ferradura", "Funil", "Farol", "Gaiola", "Garfo", "Gargantilha", "Garrafa", "Gaveta", "Guitarra", "Gaita", "Gravata", "Gorro", "Harpa", "Hélice", "Holofote", "Isqueiro", "Imã", "Isopor", "Isca", "Inalador", "Incenso", "Jarra", "Jóias", "Jaqueta", "Janela", "Lápis", "Luvas", "Lampada", "Luminária", "Leque", "Lenço", "Livro", "Lanterna", "Lixeira", "Lapiseira", "Lata", "Mola", "Maquiagem", "Meias", "Mala", "Mochila", "Navalha", "Novelo", "Navio de Brinquedo", "Óculos", "Obras de Arte", "Óculos de Sol", "Panela", "Pedra", "Papel", "Porta-Retrato", "Pen Drive", "Prato", "Pulseira", "Queijeira", "Quebra Cabeça", "Quadro", "Régua", "Ratoeira", "Roleta", "Relógio", "Remo", "Roteador", "Rímel", "Revista", "Sandália", "Sapato", "Saleiro", "Sino", "Saco", "Sacola", "Saco de pancadas", "Sabonete", "Shorts", "Tesoura", "Teclado", "Telefone", "Tampa", "Termômetro", "Trena", "Telha", "Tatame", "Tamanco", "Urna", "Ursinho de pelúcia", "Uniforme", "Umidificador de Ar", "Vassoura", "Ventilador", "Vara de pescar", "Varal", "Vaso", "Vidro", "Viseira", "Vela", "Xicara", "Xadrez", "Xale", "Xarope", "Zíper", "Zarabaana"}
	rand.Seed(time.Now().UnixNano())
	chosen := objetos[rand.Intn(len(objetos))]
	//fmt.Println(rand.Intn(len(objetos)))
	return chosen
}

func main() {
	fmt.Println("Gerando produto do Lucas Neto..")
	fmt.Println(randObject(), "do Lucas Neto")
}
