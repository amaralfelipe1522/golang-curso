package matematica

import (
	"fmt"
	"strconv"
)

//Media é bla bla bla
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}

//Este exemplo tb foi usado na aula sobre cobertura de testes
//$go test -cover
//Informa em % a cobertura de sucesso do teste
//$go test -coverprofile=resultado.out
//Exporta a cobertura de teste para um arquivo .out
//go tool cover -html=resultado.out
//Exporta e apresenta o resultado do teste em uma página HTML
