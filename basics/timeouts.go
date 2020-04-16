package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		c <- "Message sent"
	}()
	select {
	case res := <-c: //if c gives response
		{
			fmt.Println("Message received from routine ", res)
		}
	case <-time.After(time.Second * 2): //if timeout gives response
		{
			fmt.Println("Message received from Timer")
		}
	}
}
