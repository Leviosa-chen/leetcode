package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node4 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node6
	l1 := node1

	head := swapPairs(l1)
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}

}

func swapPairs1(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
