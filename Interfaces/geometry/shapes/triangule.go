package shapes

// Triangle é uma implementação da interface Shape para um Triangulo.
type Triangle struct{}

// Area é o método que define o comportamento do tipo Triangle para calcular a área.
func (t Triangle) Area(base, height int) float64 {
	return float64(base * height)
}

// Perimeter é o método que define o comportamento do tipo Triangle para calcular o perimetro.
func (t Triangle) Perimeter(base int) float64 {
	return float64(3 * base)
}
