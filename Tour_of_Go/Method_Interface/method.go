package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type myFloat float64

func (v Vertex) Hypotenuse() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Abs(v Vertex) float64 {
	return math.Abs(v.x)
}

func (f myFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Scale(v *Vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Hypotenuse())
	f := myFloat(math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(1.53)
	fmt.Println(v.Hypotenuse())
	Scale(&v, 1/1.53)
	fmt.Println(v.Hypotenuse())
}
