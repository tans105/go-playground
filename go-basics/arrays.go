package main

import "fmt"

func main() {
	fmt.Println("Hello WOrld")

	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	var arr [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = i + j
		}
	}

	fmt.Println(arr)

	strArray := []string{"a", "b", "c", "d"}
	fmt.Println(strArray)
}
