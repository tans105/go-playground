package main

import "fmt"

/*
https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3320/

First Unique Character in a String

Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
*/

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	found := make([]bool, 26)
	dictionary := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charIndex := s[i] - 'a'
		char := s[i]

		if !found[charIndex] {
			if _, ok := dictionary[char]; !ok {
				dictionary[char] = i
			} else {
				found[charIndex] = true
				dictionary[char] = -1
			}
		}
	}

	for k, v := range s {
		if dictionary[byte(v)] != -1 {
			return k
		}
	}

	return -1
}
