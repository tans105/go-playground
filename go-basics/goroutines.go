package main

import (
	"fmt"
	"strconv"
	"time"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s + ":" + strconv.Itoa(i))
	}
}

func display(str string) {
	for w := 0; w < 6; w++ {
		fmt.Println(str)
	}
}

func main() {
	f("direct")

	go f("routine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	go display("Welcome")


	/*
	when a new Goroutine executed, the Goroutine call return immediately.
	The control does not wait for Goroutine to complete their execution just like normal function they always move forward
	to the next line after the Goroutine call and ignores the value returned by the Goroutine.

	Hence we need some timeout to wait
	 */
	time.Sleep(time.Second)
	fmt.Println("Done")
}


/*
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done
 */


