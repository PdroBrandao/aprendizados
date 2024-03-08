package main

import (
	"fmt"

	"geometry/logger"
	"geometry/shapes"
)

func main() {
	log := logger.Logger{}
	log.Print("-- Geometry Shapes --")

	x, y, err := receivedParameters(log)
	if err != nil {
		log.Print(err.Error())
		return
	}
	printResults(log, x, y)
}

// receivedParameters lê os valores de x e y e retorna os mesmos.
func receivedParameters(log logger.Logger) (int, int, error) {
	var x, y int
	for {
		log.Print("Insira o valor inteiro para x (Por favor, insira um valor válido)")
		if _, err := fmt.Scan(&x); err == nil {
			break
		}
		fmt.Scanln()
	}
	for {
		log.Print("Insira o valor inteiro de y (Por favor, insira um valor válido)")
		if _, err := fmt.Scan(&y); err == nil {
			break
		}
		fmt.Scanln()
	}
	return x, y, nil
}

// printResults imprime os resultados das operações de área e perímetro de cada forma.
func printResults(log logger.Logger, x, y int) {
	shapeResultLogger := logger.ResultLogger{}
	log.Print("Abaixo os resultados:")
	shapeResultLogger.Print("RECTANGLE", shapes.Rectangle{}.Area(x, y), shapes.Rectangle{}.Perimeter(x, y))
	shapeResultLogger.Print("SQUARE", shapes.Square{}.Area(x), shapes.Square{}.Perimeter(x))
	shapeResultLogger.Print("TRIANGLE", shapes.Triangle{}.Area(x, y), shapes.Triangle{}.Perimeter(x))
	shapeResultLogger.Print("CIRCLE", shapes.Circle{}.Area(x), shapes.Circle{}.Circumference(x))
}
