package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	headA := node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	head := rotateRight(headA, 5)
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}

}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	last := head
	l := 1
	for last.Next != nil {
		last = last.Next
		l++
	}

	last.Next = head

	for i := 1; i < (l - k%l); i++ {
		head = head.Next
	}
	newhead := head.Next
	head.Next = nil
	return newhead
}
