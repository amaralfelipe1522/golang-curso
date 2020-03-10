package matematica

// No terminal é possível executar o teste com o comando: go test -v
// Ou em um diretório pai: go test ./...
import (
	"fmt"
	"testing"
)

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."
const sucessoPadrao = "Valor esperado foi retornado."

func TestMedia(t *testing.T) {
	valorEsperado := 7.29
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	} else {
		fmt.Println(sucessoPadrao)
	}
}
