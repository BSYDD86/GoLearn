package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//   1
	// 2   3
	//4 5 6 7
	//morris
	//abdbeacfcg
	nodeD := &TreeNode{4, nil, nil}
	nodeE := &TreeNode{5, nil, nil}
	nodeF := &TreeNode{6, nil, nil}
	nodeG := &TreeNode{7, nil, nil}
	nodeB := &TreeNode{2, nodeD, nodeE}
	nodeC := &TreeNode{3, nodeF, nodeG}
	nodeA := &TreeNode{1, nodeB, nodeC}
	res = []int{}
	morrisPreOrderTraversal(nodeA)
	fmt.Println(res)
}

var (
	res []int
)

func morrisPreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var morrisRight *TreeNode
	for cur != nil {
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				res = append(res, cur.Val)
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
			}
		} else {
			res = append(res, cur.Val)
		}
		cur = cur.Right
	}
}

func morrisPreTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	cur := root
	var morrisRight *TreeNode
	for cur != nil {
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				res = append(res, cur.Val)
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
			}
		} else {
			res = append(res, cur.Val)
		}
		cur = cur.Right
	}
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
