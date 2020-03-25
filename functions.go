package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, )) // 3
	fmt.Println(fibonacci(5)) // 5
	fmt.Println(myName()) //1 Tanmay
}

func sum(a int, b int) int {
	return a + b
}

func fibonacci(n int) int {
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[n-1]
}

func myName() (string, string) {
	myObject := make(map[string]string)
	myObject["name"] = "Tanmay"
	myObject["id"] = "1"
	return myObject["id"], myObject["name"]
}
