package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go arraySum(s[:len(s)/2], c) //channel to be passed in the function
	go arraySum(s[len(s)/2:], c)

	x := <-c // passing the channel 1 value into x
	y := <-c // passing the channel 2 value into y
	fmt.Println(x, y)
	time.Sleep(time.Second)
}

func arraySum(s []int, c chan int) {
	var sum int
	for _, v := range s {
		sum += v
	}

	c <- sum //assigning the sum into the channel
}
