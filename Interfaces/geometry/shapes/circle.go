package shapes

import (
	"math"
)

// Circle é uma implementação da interface Shape para um circulo.
type Circle struct{}

// Area é o método que define o comportamento do tipo Circle para calcular a área.
func (c Circle) Area(radio int) float64 {
	return float64(radio*radio) * math.Pi
}

// Circumference é o método que define o comportamento do tipo Circle para calcular a circunferência.
func (c Circle) Circumference(radio int) float64 {
	return float64(2*radio) * math.Pi
}
