package main

import (
	"fmt"
	stack "go-playground/go-datastructure/collections"
)

/*
Remove K Digits

Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:

    The length of num is less than 10002 and will be ≥ k.
    The given num does not contain any leading zero.

Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

Example 2:

Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

Example 3:

Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.

*/

func main() {
	removeKdigits("1432219", 3)
}

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}

	if k == len(num) {
		return "0"
	}

	stk := stack.New()

	for _, c := range num {
		for !stk.IsEmpty() && k > 0 && stk.Peek().(int32) > c {
			stk.Pop()
			k--
		}

		stk.Push(c)
	}

	for i := 0; i < k; i++ {
		stk.Pop()
	}

	var sb []int32

	for !stk.IsEmpty() {
		sb = append(sb, stk.Pop().(int32))
	}

	sb = reverse(sb)

	fmt.Println(sb)

	return ""
}

func reverse(numbers []int32) []int32 {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

