package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 1000; i++ {
		z -= (math.Pow(z, 2.0) - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
