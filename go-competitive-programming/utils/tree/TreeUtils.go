package tree

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func Traverse(node *Node) {
	if node == nil {
		return
	}

	Traverse(node.Left)
	fmt.Print(node.Val, "\t")
	Traverse(node.Right)
}
