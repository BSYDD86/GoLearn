package main

import "fmt"

func main() {
	root := &TreeNode{1, nil, nil}
	order := levelOrder(root)
	fmt.Println(order)
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	temp := []*TreeNode{root}
	for len(temp) > 0 {
		n := len(temp)
		path := make([]int, n)
		for i := 0; i < n; i++ {
			if temp[i].Left != nil {
				temp = append(temp, temp[i].Left)
			}
			if temp[i].Right != nil {
				temp = append(temp, temp[i].Right)
			}
			path[i] = temp[i].Val
		}
		res = append(res, path)
		temp = temp[n:]
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
