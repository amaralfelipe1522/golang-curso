package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - indices: esperado (%d) <> encontrado (%d)"

func TestIndex(t *testing.T) {
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Atumalaka HAHAHA", "Atumalaka", 0}, //Existe no indice 0
		{"", "", 0},                          //Existe no indice 0
		{"Xablau", "xablau", -1},             //Não existe
		{"Picles", "i", 1},                   //Letra i no indice 1
	}

	for _, teste := range testes {
		t.Logf("Massa de testes: %v", teste)
		testeAtual := strings.Index(teste.texto, teste.parte)

		if testeAtual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, testeAtual)
		}
	}
}
