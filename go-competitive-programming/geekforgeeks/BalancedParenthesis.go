package main

import "fmt"
import "github.com/golang-collections/collections/stack"

/*
https://www.geeksforgeeks.org/check-for-balanced-parentheses-in-an-expression/

Check for balanced parentheses in an expression

Given an expression string exp , write a program to examine whether the pairs and the orders of “{“,”}”,”(“,”)”,”[“,”]” are correct in exp.

Example:

Input: exp = “[()]{}{[()()]()}”
Output: Balanced



Input: exp = “[(])”
Output: Not Balanced
*/
func main() {
	fmt.Println(isBalanced("[()]{}{[()()]()}"))
	fmt.Println(isBalanced("[()]"))
	fmt.Println(isBalanced("[()"))
}

func isBalanced(expr string) bool {
	stack := stack.New()
	index := 0
	isBalanced := true
	for index < len(expr) {
		temp := string(expr[index])

		if temp == "[" || temp == "(" || temp == "{" {
			stack.Push(temp)
		} else {
			peek := stack.Peek()
			opp := getOpposite(temp)
			if peek == opp {
				stack.Pop()
			} else {
				isBalanced = false
				break
			}
		}
		index++
	}

	return stack.Len() == 0  && isBalanced
}

func getOpposite(peek interface{}) string {
	switch peek {
	case "]":
		return "["
	case ")":
		return "("
	default:
		return "{"
	}

}
