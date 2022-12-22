package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	tri := triangle{height: 5, base: 3}
	squ := square{4.3}

	printArea(tri)
	printArea(squ)
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
