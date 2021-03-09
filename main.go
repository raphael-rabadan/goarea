package goarea

import "math"

// Pi é uma proporção numérico definida pela relção entre
// o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Circ calcula a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a área de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// TrianguloEq Fuuncioonaa
func TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
