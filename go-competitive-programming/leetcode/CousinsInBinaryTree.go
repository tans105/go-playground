package main

import (
	"fmt"
	"go-playground/go-competitive-programming/utils/tree"
)

/*
https://leetcode.com/explore/challenge/card/may-leetcoding-challenge/534/week-1-may-1st-may-7th/3322/

Cousins in Binary Tree
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.
Two nodes of a binary tree are cousins if they have the same depth, but have different parents.
We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.
Return true if and only if the nodes corresponding to the values x and y are cousins.

Example 1:

Input: root = [1,2,3,4], x = 4, y = 3
Output: false

Example 2:

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true

Example 3:

Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false

Note:

    The number of nodes in the tree will be between 2 and 100.
    Each node has a unique integer value from 1 to 100.

*/

func main() {
	root := tree.Node{1,
		&tree.Node{2,
			nil,
			&tree.Node{4,
				nil,
				nil}},
		&tree.Node{3,
			nil,
			&tree.Node{5,
				nil,
				nil}}}
	
	fmt.Println(isCousins(&root, 4, 3))
}


func isCousins(root *tree.Node, x int, y int) bool {
	var xData TreeData
	var yData TreeData

	xData = isCousinsHelper(root, nil, x, 0)
	yData = isCousinsHelper(root, nil, y, 0)

	return xData.height == yData.height && xData.parent != yData.parent
}

func isCousinsHelper(root *tree.Node, parent *tree.Node, val int, heightSoFar int) TreeData {

	if root == nil {
		return TreeData{}
	}

	if root.Val == val {
		if parent != nil {
			return TreeData{
				height: heightSoFar,
				parent: parent.Val,
			}
		} else {
			return TreeData{}
		}
	}

	data := isCousinsHelper(root.Left, root, val, heightSoFar+1)
	if data != (TreeData{}) {
		return data
	}
	data = isCousinsHelper(root.Right ,root, val, heightSoFar+1)
	return data
}

type TreeData struct {
	height int
	parent int
}
