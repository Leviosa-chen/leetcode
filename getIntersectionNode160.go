package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := &ListNode{Val: 4}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 8}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 5}
	node7 := &ListNode{Val: 6}
	node8 := &ListNode{Val: 1}

	headA := node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	headB := node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node3

	head := getIntersectionNode(headA, headB)
	// 输出
	println(head.Val)

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	hA, hB := headA, headB
	for hA != hB {
		if hA == nil {
			hA = headB
		} else {
			hA = hA.Next
		}
		if hB == nil {
			hB = headA
		} else {
			hB = hB.Next
		}
	}
	return hA
}
