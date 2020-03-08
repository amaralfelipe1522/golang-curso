package main

import "math"

// Ponto representa uma coordenada no plano cartesiano
// Maiúscula significa ser pública dentro e fora do pacote, do contrário é privado APENAS dentro do pacote (ou iniciando com underline: _Example)
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return //retorno já nomeado
}

// Distancia é responsável por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
