package main

import "fmt"

/*
https://www.geeksforgeeks.org/check-if-an-array-can-be-divided-into-pairs-whose-sum-is-divisible-by-k/

Check if an array can be divided into pairs whose sum is divisible by k
Given an array of integers and a number k, write a function that returns true if given array can be divided into pairs such that sum of every pair is divisible by k.
Examples:

Input: arr[] = {9, 7, 5, 3},
k = 6
Output: True
We can divide array into (9, 3) and
(7, 5). Sum of both of these pairs
is a multiple of 6.



Input: arr[] = {92, 75, 65, 48, 45, 35},
k = 10
Output: True
We can divide array into (92, 48), (75, 65)
and (45, 35). Sum of all these pairs is a
multiple of 10.

Input: arr[] = {91, 74, 66, 48}, k = 10
Output: False
 */
func main() {
	fmt.Println(isDivisible([]int{92, 75, 65, 48, 45, 35}, 10))
	fmt.Println(isDivisible([]int{92, 75, 65, 48}, 9))
	fmt.Println(isDivisible([]int{9, 7, 5, 3}, 6))
	fmt.Println(isDivisible([]int{91, 74, 66, 48}, 10))
}

func isDivisible(nums []int, k int) bool {
	remainderMap := make(map[int]int)

	for _, num := range nums {
		count, present := remainderMap[num%k]
		if present {
			remainderMap[num%k] = count + 1
		} else {
			remainderMap[num%k] = 1
		}
	}

	for key, value := range remainderMap {
		if key == 0 {
			if value%2 != 0 {
				return false
			}
		} else if (key*value)%k == 0 {
			continue
		} else if value == remainderMap[k-key] {
			continue
		} else {
			return false
		}
	}
	return true
}
