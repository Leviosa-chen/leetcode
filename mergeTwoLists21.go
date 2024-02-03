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

	head := mergeTwoLists(l1, l2)
	// 输出
	current := head
	for current != nil {
		current = current.Next
	}
}
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode

	// 比较 替换 使 list的第一个数小于第二个
	if list1.Val > list2.Val {
		head = list2
		list2 = list1
		list1 = head
	} else {
		head = list1
	}

	list1 = list1.Next
	current := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}
