package main

type ListNode struct {
	Val        int
	Prev, Next *ListNode
}

func main() {
	dummy := &ListNode{}
	println(dummy.Val)
}
