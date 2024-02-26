package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type bot interface {
	getArea() float64
}

func main() {
	t := triangle{10, 8}
	s := square{10}
	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(b bot) {
	fmt.Println(b.getArea())
}
