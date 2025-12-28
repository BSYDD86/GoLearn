package main

import "fmt"

func main() {
	vals := []int{1, 2, 3, 4, 5, 6, 7}
	root := buildTree(vals)
	to_delete := []int{3, 5}
	nodes := delNodes(root, to_delete)
	fmt.Println(nodes)
}

func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}

	for i := 0; i < len(vals); i++ {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		if leftIndex < len(vals) {
			nodes[i].Left = nodes[leftIndex]
		}
		if rightIndex < len(vals) {
			nodes[i].Right = nodes[rightIndex]
		}
	}

	return nodes[0] // 根节点
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	res []*TreeNode
	m   map[int]bool
)

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res = []*TreeNode{}
	m = map[int]bool{}
	for _, v := range to_delete {
		m[v] = true
	}
	dfs(root, true)
	return res
}

func dfs(root *TreeNode, deleted bool) bool {
	if root == nil {
		return true
	}
	curDeleted := false
	if m[root.Val] {
		curDeleted = true
	}
	if deleted && !curDeleted {
		res = append(res, root)
	}
	l := dfs(root.Left, curDeleted)
	r := dfs(root.Right, curDeleted)
	if l {
		root.Left = nil
	}
	if r {
		root.Right = nil
	}
	return curDeleted
}
