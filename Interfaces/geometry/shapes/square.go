package shapes

// Square é uma implementação da interface Shape para um quadrado.
type Square struct{}

func (s Square) Area(side int) float64 {
	return float64(side * side)
}

func (s Square) Perimeter(side int) float64 {
	return float64(4 * side)
}
