package main

import "fmt"

func main() {
	nextInt := getSeq()
	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	nextInt = getSeq()
	fmt.Println(nextInt()) //1
}

func getSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
