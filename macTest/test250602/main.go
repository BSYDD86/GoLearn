package main

import "fmt"

func main() {
	node1 := &TreeNode{1, nil, nil}
	node0 := &TreeNode{0, node1, nil}
	root := &TreeNode{2, nil, node0}
	fmt.Println(maxAncestorDiff(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res int
)

func maxAncestorDiff(root *TreeNode) int {
	res = 0
	if root.Left != nil {
		dfs(root.Left, root.Val, root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, root.Val, root.Val)
	}

	return res
}
func dfs(root *TreeNode, iMin int, iMax int) {
	l := root.Left
	r := root.Right
	if l == r {
		temp1 := root.Val - iMin
		temp2 := root.Val - iMax
		if temp1 < 0 {
			temp1 = -temp1
		}
		if temp2 < 0 {
			temp2 = -temp2
		}
		res = max(res, temp1, temp2)
		return
	}
	if l != nil {
		dfs(root.Left, min(iMin, root.Val), max(iMax, root.Val))
	}
	if r != nil {
		dfs(root.Right, min(iMin, root.Val), max(iMax, root.Val))
	}
}
