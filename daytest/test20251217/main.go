package main

import (
	"fmt"
	"strings"
)

func main() {
	//a := true
	//fmt.Println(a)
	//if a {
	//	a = !a
	//}
	//fmt.Println(a)

	str := "a,b"
	r := strings.Split(str, ",")
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		n := len(queue)
		path := []int{}
		for _, v := range queue {
			path = append(path, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		if reverse {
			l := 0
			r := len(path)
			for l < r {
				path[l], path[r] = path[r], path[l]
				l++
				r--
			}
			reverse = false
		} else {
			reverse = true
		}
		queue = queue[n:]
		res = append(res, path)
	}
	return res
}
