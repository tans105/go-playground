package main

import (
	"errors"
	"fmt"
)

//By convention, errors are the last return value and have type error, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42") //errors.New constructs a basic error value with the given error message.
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

/*
It’s possible to use custom types as errors by implementing the Error() method on them.
Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.
*/
func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e) //f1 failed: can't work with 42
		} else {
			fmt.Println("f1 worked:", r) //f1 worked: 10
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e) //f2 failed: 42 - can't work with it
		} else {
			fmt.Println("f2 worked:", r) //f2 worked: 10
		}
	}

	_, e := f2(42)
	ae, ok := e.(argError)

	if ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
