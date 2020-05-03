package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := make([]map[string]string, 3)

	map1 := map[string]string{"name": "User1", "id": "1"}
	map2 := map[string]string{"name": "User2", "id": "2"}
	map3 := map[string]string{"name": "User3", "id": "3"}

	arr[0] = map1
	arr[1] = map2
	arr[2] = map3

	fmt.Println(arr) // [map[id:1 name:User1] map[id:2 name:User2] map[id:3 name:User3]]

	//dynamic creation of list of maps
	for i := 0; i < 3; i++ {
		id := strconv.Itoa(i)
		arr[i] = map[string]string{"name": "User" + id, "id": id}
	}

	fmt.Println(arr) // [map[id:0 name:User0] map[id:1 name:User1] map[id:2 name:User2]]
}
