package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 0, 1, 0, 1, 0, 1}
	root := buidTreeByArr(arr, 0)

	fmt.Println(sumRootToLeaf(root))
}

func buidTreeByArr(path []int, idx int) *TreeNode {
	if idx >= len(path) {
		return nil
	}
	root := &TreeNode{path[idx], buidTreeByArr(path, 2*idx+1), buidTreeByArr(path, 2*idx+2)}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, []int{})
}
func dfs(root *TreeNode, path []int) int {
	l := root.Left
	r := root.Right
	var res int
	path = append(path, root.Val)
	if l == nil && r == nil {
		for i := 0; i < len(path); i++ {
			res = res*2 + path[i]
		}
		return res
	} else if l == nil || r == nil {

		if l == nil {
			return dfs(root.Right, path)
		} else if r == nil {
			return dfs(root.Left, path)
		}
	} else {

		return dfs(root.Left, path) + dfs(root.Right, path)
	}
	return 0
}
