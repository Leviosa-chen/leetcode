package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var root *TreeNode
	root = insertLevelOrder(arr, root, 0, len(arr))

	fmt.Println("Inorder traversal of the constructed binary tree:")
	printInOrder(root)
	fmt.Println("depth:")
	println(maxDepth(root))
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func insertLevelOrder(arr []int, root *TreeNode, i int, n int) *TreeNode {
	if i < n {
		temp := &TreeNode{Val: arr[i]}
		root = temp

		root.Left = insertLevelOrder(arr, root.Left, 2*i+1, n)
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2, n)
	}
	return root
}

func printInOrder(root *TreeNode) {
	if root != nil {
		printInOrder(root.Left)
		printInOrder(root.Right)
		fmt.Printf("%d ", root.Val)
	}
}

func maxDepth(root *TreeNode) int {
	return bfs(root)
}

func bfs(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		for l > 0 {
			node := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			l--
		}

		depth++
	}
	return depth
}
