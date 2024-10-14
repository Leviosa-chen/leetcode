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
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node4

	head := node1

	node := hasCircle(head)
	if node != nil {
		println(node.Val)
	} else {
		println("no")
	}

	// 输出
	//current := head
	//for current != nil {
	//	println(current.Val)
	//	current = current.Next
	//}

}

func hasCircle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head

	has := false
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			has = true
			break
		}
	}
	if has {
		slow = head
		for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next
			if fast == slow {
				return fast
			}
		}
	}

	return nil
}
