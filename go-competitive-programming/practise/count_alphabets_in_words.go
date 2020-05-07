package main

import "fmt"

func main() {
	countMap := make(map[string]int)
	word := "tanmay"

	for _, c := range word {
		char := string(c) //convert char ascii to character

		if _, present := countMap[char]; present { // countMap[char] returns 2 values, 1st is count and 2nd is bool if key is present or not
			countMap[char]++
		} else {
			countMap[char] = 1
		}
	}
	fmt.Println(countMap) //map[a:2 m:1 n:1 t:1 y:1]
}
