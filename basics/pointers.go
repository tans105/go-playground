package main

import "fmt"

/*
* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.

& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
 */
func main() {
	i := 1
	fmt.Println(i) //1

	zeroVal(i)
	fmt.Println(i) //1

	zeroPtr(&i)
	fmt.Println(i)  //0
	fmt.Println(&i) //0xc000014080
}

func zeroVal(val int) { //pass by value
	val = 0
}

func zeroPtr(val *int) { //pass by reference
	*val = 0
}
