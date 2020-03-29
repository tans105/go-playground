package main

import "fmt"

type Person struct {
	fname   string
	favFood []string
}

func (p Person) walk() string {
	return p.fname + " is walking"
}

func main() {
	p := Person{
		"Tanmay",
		[]string{"a", "b", "c", "d"},
	}

	fmt.Println(p)

	for _, v := range p.favFood {
		fmt.Println(v)
	}

	fmt.Println(p.walk())
}
