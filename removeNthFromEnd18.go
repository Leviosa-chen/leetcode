package main

func main() {

	node1 := &ListNode{Val: 1}
	//node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}
	//node4 := &ListNode{Val: 4}
	//node5 := &ListNode{Val: 5}
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4
	//node4.Next = node5

	l1 := node1

	head := removeNthFromEnd(l1, 1)
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}

	// 删除第一个
	if fast == nil || fast.Next == nil {
		head = head.Next
		return head
	}

	pre := slow
	for fast != nil && fast.Next != nil {
		pre = slow
		fast = fast.Next
		slow = slow.Next
	}

	pre.Next = slow.Next
	return head
}
