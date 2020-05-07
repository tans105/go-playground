package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	var sec int = 0
	var min int = 0
	var hour int = 0
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C: //ticker start
				sec++
				if sec == 60 {
					sec = 0
					min++
				}
				if min == 60 {
					min = 0
					hour++
				}
				if hour == 12 {
					hour = 0
				}
				fmt.Println(hour, min, sec)
			}
		}
	}()

	time.Sleep(16000 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
