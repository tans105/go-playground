package main

import "fmt"

func main() {

	var i int = 0
	for {
		if i == 100 {
			break
		}
		switch {
		case i <= 10:
			{
				fmt.Println("Its less than 10 :", i)
				break
			}
		case i > 10 && i < 90:
			{
				fmt.Println("Its more than 10 :", i)
				break
			}
		case i > 90:
			{
				fmt.Println("Its more than 90 :", i)
				break
			}
		}

		i++
	}

}
