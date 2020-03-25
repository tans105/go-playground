package main

import (
	"fmt"
	"reflect"
	)

func main() {
	fmt.Println("hello world")
	fmt.Println("1 + 1 = ", 1+1, " its working")

	var a string = "2"
	fmt.Println(a)

	var b bool = true
	fmt.Println(b)

	var name string = "Tanamy"
	fmt.Println(name)

	f := "apple"
	fmt.Println(f)

	isTanmay := true
	fmt.Println(reflect.TypeOf(isTanmay))

	if isTanmay {
		fmt.Println(true)
	}
}
