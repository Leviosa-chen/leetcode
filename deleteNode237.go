package main

func main() {

	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 5}
	node3 := &ListNode{Val: 1}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 9}

	headA := node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	deleteNode(node2)
	// 输出
	current := headA
	for current != nil {
		println(current.Val)
		current = current.Next
	}

}

func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
			break
		}
		node = node.Next
	}
}
