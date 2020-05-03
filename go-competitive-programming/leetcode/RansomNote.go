package main

import "fmt"

/*
https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3318/

Ransom Note

Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

func canConstruct(ransomNote string, magazine string) bool {
	ransomMap := make([]int, 26)

	for i := 0; i < len(ransomNote); i++ {
		ransomMap[ransomNote[i]-'a']++
	}

	for i := 0; i < len(magazine); i++ {
		if ransomMap[magazine[i]-'a'] > 0 {
			ransomMap[magazine[i]-'a']--
		}
	}

	for i := 0; i < len(ransomMap); i++ {
		if ransomMap[i] > 0 {
			return false
		}
	}

	return true
}
