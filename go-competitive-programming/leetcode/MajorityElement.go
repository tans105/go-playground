package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{1}))
}

func majorityElement(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}

	Map := make(map[int]int)
	length := len(nums) / 2

	for _, num := range nums {
		count, present := Map[num]
		if present {
			if count+1 > length {
				return num
			}
			Map[num] = count + 1
		} else {
			Map[num] = 1
		}
	}


	return -1
}
