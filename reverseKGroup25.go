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

	head := reverseKGroup(l1, 3)
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}

}
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{0, head}
	temp := dummyHead
	for temp.Next != nil {
		current := temp.Next
		for i := 0; i < k-1 && current != nil && current.Next != nil; i++ {
			current = current.Next
		}
		nextHead := current.Next
		current.Next = nil
		h, e := reverseList2(temp.Next)
		temp.Next = h
		e.Next = nextHead
		temp = e
	}
	return dummyHead.Next
}

func reverseList2(head *ListNode) (*ListNode, *ListNode) {
	current := head
	var pre *ListNode
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}

	return pre, head
}
