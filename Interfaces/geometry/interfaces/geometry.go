package interfaces

// Shape é uma interface que define métodos para calcular propriedades geométricas.
type Shape interface {
	// Area calcula a área da forma.
	// Retorna a área da forma como um valor float64.
	Area(n ...int) float64
	// Perimeter calcula o perimetro da forma.
	// Retorna a perimetro da forma como um valor float64.
	Perimeter(n ...int) float64
	// Circumference calcula a circunferência da forma.
	// Retorna a área da forma como um valor float64.
	Circumference(n ...int) float64
}
