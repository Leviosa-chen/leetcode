package main

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3

	l1 := node1

	nodea := &ListNode{Val: 1}
	nodeb := &ListNode{Val: 3}
	nodec := &ListNode{Val: 4}
	nodea.Next = nodeb
	nodeb.Next = nodec

	l2 := nodea

	head := mergeKLists([]*ListNode{l1, l2})
	// 输出
	current := head
	for current != nil {
		println(current.Val)
		current = current.Next
	}

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var head *ListNode
	for i := 0; i < len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}
