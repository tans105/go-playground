package main

import "fmt"

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.length * s.length
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

func getShapeArea(s shape) float64 {
	return s.area()
}

func main() {
	s := Square{10}
	c := Circle{10}
	fmt.Println(getShapeArea(s)) //100
	fmt.Println(getShapeArea(c)) //314
}
