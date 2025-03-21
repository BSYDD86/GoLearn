package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1}
	fmt.Println(isPalindrome(createList(arr)))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(arr []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arr {
		t := &ListNode{Val: v}
		cur.Next = t
		cur = t
	}
	return dummy.Next
}

func isPalindrome(head *ListNode) bool {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil {
		slow = slow.Next
	}
	headReverse := reverseListNode(slow)
	slow = head
	for headReverse != nil {
		if slow.Val != headReverse.Val {
			return false
		}
		slow = slow.Next
		headReverse = headReverse.Next
	}
	return true

}
func reverseListNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	return pre
}
