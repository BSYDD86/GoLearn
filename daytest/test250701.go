package main

import "fmt"

func main() {
	arr := []int{5, 6, 6, -1, -1, 8, 6, 10, -1, 5}
	tree := buildTree(arr)
	fmt.Println(tree)
	fmt.Println(tree)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	idx := 1

	for len(queue) > 0 && idx < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 左子节点
		if idx < len(arr) && arr[idx] != -1 {
			node.Left = &TreeNode{Val: arr[idx]}
			queue = append(queue, node.Left)
		}
		idx++

		// 右子节点
		if idx < len(arr) && arr[idx] != -1 {
			node.Right = &TreeNode{Val: arr[idx]}
			queue = append(queue, node.Right)
		}
		idx++
	}

	return root
}
