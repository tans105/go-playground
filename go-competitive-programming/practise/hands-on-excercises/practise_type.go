package main

import "fmt"

type gator int
type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am a flamingo")
}

func bayou(creature swampCreature) {
	creature.greeting()
}

func main() {
	var g1 gator
	fmt.Println(g1)                 //0
	fmt.Printf("Type =>  %T\n", g1) //Type =>  main.gator0

	var x int
	fmt.Println(x)        //0
	fmt.Printf("%T\n", x) //int

	g1 = gator(x) //you cannot directly assign, type is different. It needs casting
	//x = g1 - wrong
	//x = int(g1) - right

	g1.greeting() //Hello, I am a gator

	var f1 flamingo

	bayou(g1) //Hello, I am a gator
	bayou(f1) //Hello, I am a flamingo
}
