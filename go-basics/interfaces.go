package main

import "fmt"

func main() {
	arr := []Animal{Dog{"Doggy"}, Cat{"Kitty"}, Human{"Insaan"}}
	for _, v := range arr {
		fmt.Println(v, v.speak())
	}
}

/*
We define an Animal as being any type that has a method named Speak.
The Speak method takes no arguments and returns a string.
Any type that defines this method is said to satisfy the Animal interface.
There is no implements keyword in Go; whether or not a type satisfies an interface is determined automatically.
*/
type Animal interface {
	speak() string
}

type Dog struct {
	name string
}

type Human struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) speak() string {
	return "woof"
}

func (c Cat) speak() string {
	return "meow"
}

func (c Human) speak() string {
	return "grr"
}
