package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool, 1)
	fmt.Println("Starting operation")
	worker(c)
	done := <-c
	//Will remain at pause till worker responds.
	fmt.Println("Its completed..", done)
}

func worker(c chan bool) {
	fmt.Println("Starting doing a task!")
	time.Sleep(time.Second * 5) // operatiom taking 2 second
	c <- true
}
