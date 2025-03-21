package main

import "fmt"

func main() {
	lists := []*ListNode{&ListNode{Val: 1},
		&ListNode{Val: 1},
		&ListNode{Val: 3},
		&ListNode{Val: 4}}
	fmt.Println(mergeKLists(lists))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	for k > 0 {
		idx := 0
		for i := 0; i < k; i += 2 {
			if i == k-1 {
				lists[idx] = lists[i]
			} else {
				lists[idx] = mergeTwoLists(lists[i], lists[i+1])
			}
		}
		k = idx
	}
	return lists[0]
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
		if list1 != nil {
			cur.Next = list1
		} else {
			cur.Next = list2
		}
	}
	return dummy.Next
}
