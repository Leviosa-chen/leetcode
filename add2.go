package main

func main() {

	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3

	l1 := node1

	nodea := &ListNode{Val: 5}
	nodeb := &ListNode{Val: 6}
	nodec := &ListNode{Val: 4}
	nodea.Next = nodeb
	nodeb.Next = nodec

	l2 := nodea

	head := addTwoNumbers2(l1, l2)
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	var head *ListNode
	pre := head
	for l1 != nil || l2 != nil || add != 0 {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
		}
		sum := l1Val + l2Val + add
		current := ListNode{}
		current.Val = sum % 10
		add = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if pre == nil {
			head = &current
			pre = &current
		} else {
			pre.Next = &current
			pre = pre.Next
		}
	}
	return head
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	head := ListNode{Val: sum % 10}
	head.Next = addTwoNumbers2(l1.Next, l2.Next)
	if sum > 9 {
		head.Next = addTwoNumbers2(head.Next, &ListNode{Val: 1})
	}

	return &head
}
