package logger

import "fmt"

// Logger é a implementação da interface Log.
type Logger struct{}

// Print é o método que define o comportamento do tipo Logger.
func (l Logger) Print(msg string) {
	fmt.Println(msg)
}

// ResultLogger é a implementação da interface Log.
type ResultLogger struct{}

// Print é o método que define o comportamento do tipo ResultLogger.
func (il ResultLogger) Print(shape string, area, perimeter float64) {
	fmt.Printf("[RESULT] %s - Area = %f -- Per = %f\n", shape, area, perimeter)
}
