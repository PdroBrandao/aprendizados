package shapes

// Rectangle é uma implementação da interface Shape para um Retangulo.
type Rectangle struct{}

// Area é o método que define o comportamento do tipo Rectangle para calcular a área.
func (r Rectangle) Area(length, width int) float64 {
	return float64(length * width)
}

// Perimeter é o método que define o comportamento do tipo Rectangle para calcular o perimetro.
func (r Rectangle) Perimeter(length, width int) float64 {
	return float64(2 * (length + width))
}
