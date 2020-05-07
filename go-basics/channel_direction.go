package main

import "fmt"

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	m1(c1, "Hi Tanmay...")
	m2(c1, c2)
	msg := <-c2
	fmt.Println("Message recieved: ", msg)

}

func m1(c chan<- string, message string) { //1st route
	//see the param syntax, c will only recieve
	c <- message
}

func m2(c1 <-chan string, c2 chan<- string) { //2nd route
	//see the param syntax, c1 will send, c2 will receive
	msg := <-c1
	c2 <- msg
}
