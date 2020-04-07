package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr) // [1 2 3]

	for index, value := range arr {
		fmt.Println(index, value)
		/*
			0 1
			1 2
			2 3
		*/
	}

	myMap := make(map[string]string)
	//myMap := map[string]string{"foo": "Tanmay", "bar": "awasthi"}

	myMap["foo"] = "Tanmay"
	myMap["bar"] = "Awasthi"

	for key, value := range myMap {
		fmt.Println(key, value)
		/*
			foo Tanmay
			Bar Awasthi
		*/
	}

	for _, b := range "tanmay" {
		fmt.Println(b - 'a')
		/*
		19
		0
		13
		12
		0
		24
		 */
	}

}
