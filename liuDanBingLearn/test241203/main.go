package main

import "fmt"

type Node struct {
	Next *Node
	Val  int
}

func main() {
	node1 := Node{Val: 1}
	node2 := Node{Val: 2}
	node3 := Node{Val: 3}
	node4 := Node{Val: 4}
	node1.Next = &node2
	node2.Next = &node3
	node3.Next = &node4

	cur := &node1
	for cur != nil {
		fmt.Printf("%+v\n", cur.Val)

		cur = cur.Next
	}

}
