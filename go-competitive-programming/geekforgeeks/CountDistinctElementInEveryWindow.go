package main

import "fmt"

/*
https://practice.geeksforgeeks.org/problems/count-distinct-elements-in-every-window/1

Count distinct elements in every window

Given an array A[] of size N and an integer K. Your task is to complete the function countDistinct() which prints the count of distinct numbers in all windows of size k in the array A[].

Input:
The first line of input contains an integer T denoting the number of test cases. Then T test cases follow. Each test case contains two integers N and K. Then in the next line are N space separated values of the array A[].

Output:
For each test case in a new line print the space separated values denoting counts of distinct numbers in all windows of size k in the array A[].

Constraints:
1 <= T <= 100
1 <= N <= K <= 105
1 <= A[i] <= 105 , for each valid i

Example(To be used only for expected output):
Input:
2
7 4
1 2 1 3 4 2 3
3 2
4 1 1

Output:
3 4 4 3
2 1
 */
func main() {
	countDistinct([]int{3, 3, 3}, 2)
}

func countDistinct(nums []int, k int) {
	Map := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if i < k {
			count, present := Map[nums[i]]
			if present {
				Map[nums[i]] = count + 1
			} else {
				Map[nums[i]] = 1
			}
		} else {
			fmt.Print(len(Map), "\t")
			indexToBeRemoved := i - k

			count, _ := Map[nums[indexToBeRemoved]]

			if count == 1 {
				delete(Map, nums[indexToBeRemoved])
			} else {
				Map[nums[indexToBeRemoved]] = count - 1
			}

			count, _ = Map[nums[i]]
			Map[nums[i]] = count + 1
		}
	}
	fmt.Print(len(Map))
	fmt.Println()
}
