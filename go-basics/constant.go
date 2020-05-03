package main

import "fmt"
import "math"

const name = "Tanmay outer"

func main () {

	const name = "Tanmay inner"
	fmt.Println("Test")
	fmt.Println(name)
	fmt.Println(math.Sin(30))
	main1()
}

func main1() {
	fmt.Println(name)
}
