package main

import "fmt"

func main() {
	fmt.Println(fact(5))
}

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
