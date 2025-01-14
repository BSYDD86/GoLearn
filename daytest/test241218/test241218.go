package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	add := 0
	for l1 != nil && l2 != nil {
		temp := l1.Val + l2.Val + add
		cur := &ListNode{Val: temp % 10}
		add = temp / 10
		pre.Next = cur
		pre = cur
	}
	var tempNode *ListNode
	if l1 != nil {
		tempNode = l1
	} else {
		tempNode = l2
	}
	for tempNode != nil {
		temp := tempNode.Val + add
		cur := ListNode{Val: temp % 10}
		add = temp / 10
		pre.Next = cur
		pre = cur
	}
	if add != 0 {
		pre.Next = &ListNode{Val: add}
	}
	return dummy.Next
}
