package main

import "fmt"

func main() {
	rect := rectangle{5, 10}
	rectPtr := &rect

	fmt.Println(area(rect))
	fmt.Println(perimeter(rectPtr))
}

func area(rect rectangle) int {
	return rect.length * rect.width
}

func perimeter(rect *rectangle) int {
	return 2*rect.length + 2*rect.width
}

type rectangle struct {
	width  int
	length int
}
