package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}

	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7

	root := node1
	fmt.Println(behindOrderTraversal(root))
}

/**
 * Definition for a binary tree node.

 */
func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	left := inorderTraversal(root.Left)
	result = append(result, left...)
	result = append(result, root.Val)
	right := inorderTraversal(root.Right)
	result = append(result, right...)
	return result
}

func midOrderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	left := midOrderTraversal(root.Left)
	result = append(result, left...)
	right := midOrderTraversal(root.Right)
	result = append(result, right...)
	return result
}

func behindOrderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	left := behindOrderTraversal(root.Left)
	result = append(result, left...)
	right := behindOrderTraversal(root.Right)
	result = append(result, right...)
	result = append(result, root.Val)
	return result
}
