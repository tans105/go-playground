package main

import (
	"fmt"
	"time"
)

func main() {
	//Example 1
	ch := make(chan int, 2) //imagine them as threadpools, thread gets released only when someone reads from it
	//so if we create 2 threads and pass 3 values, it will throw a deadlock error
	ch <- 1
	ch <- 2
	//ch <- 3 --> will throw a deadlock
	fmt.Println(<-ch)

	fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//Example - 2
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)

	}
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

