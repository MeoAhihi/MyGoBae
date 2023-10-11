package main

import (
	"math"
)

// Input
var x = 2.4
var n = 3

// Process
func Factorial(inputNum int) float64 {
	var result = 1
	for i := 2; i < inputNum; i++ {
		result *= i
	}
	return float64(result)
}

func NaiveMaclaurinSeries(x float64, n int) float64 {
	result := 0.0
	for i := 1; i < n; i++ {
		result += math.Pow(float64(x), float64(n)) / float64(Factorial(n))
	}
	return float64(result)
}

// Output
// func main() {
// 	fmt.Printf("Result: %6.2f", NaiveMaclaurinSeries(x, n))

// }
