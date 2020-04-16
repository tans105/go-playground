package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more { // more value will be false if the channel is closed
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs") //called when channel is closed
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		fmt.Println("sent job", j)
		jobs <- j
		time.Sleep(time.Second)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done //this will hold the execution of main thread. (Blocking execution)
}
