package main

import "fmt"

type TreeNode struct {
	Val   byte
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//   a
	// b   c
	//d e f g
	//morris
	//abdbeacfcg
	nodeD := &TreeNode{'d', nil, nil}
	nodeE := &TreeNode{'e', nil, nil}
	nodeF := &TreeNode{'f', nil, nil}
	nodeG := &TreeNode{'g', nil, nil}
	nodeB := &TreeNode{'b', nodeD, nodeE}
	nodeC := &TreeNode{'c', nodeF, nodeG}
	nodeA := &TreeNode{'a', nodeB, nodeC}
	morrisTraversal(nodeA)
}
func morrisTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var morrisRight *TreeNode
	for cur != nil {
		fmt.Println(string(cur.Val))
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
			}
		}
		cur = cur.Right
	}
}
