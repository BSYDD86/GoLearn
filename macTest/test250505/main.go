package main

import (
	"fmt"
)

func main() {
}

var (
	res int
)

func sumNumbers(root *TreeNode) int {
	res = 0
	dfs(root, 0)
	return res
}
func dfs(root *TreeNode, sum int) {
	if root.Left == nil && root.Right != nil {
		res += sum
		return
	}
	sum = sum*10 + root.Val
	if root.Left != nil {
		dfs(root.Left, sum)
	}
	if root.Right != nil {
		dfs(root.Right, sum)
	}
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return targetSum == 0
	}
	targetSum -= root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

func balancedString(s string) int {
	res := 0
	target := len(s) / 4
	m := make([]int, 4)
	pre := 0
	for i, c := range s {
		m[c-'Q']++
		for m[c-'Q'] > target {
			m[s[pre]-'Q']--
			pre++
		}
		res = max(res, i-pre+1)
	}
	return len(s) - res
}
func serialize(node *TreeNode) string {
	res := []byte{}
	if node == nil {
		return string(res)
	}
	st := []*TreeNode{node}
	for len(st) > 0 {
		tempSize := len(st)
		for i := 0; i < tempSize; i++ {
			if st[i] == nil {
				res = append(res, ',')
			} else {
				res = append(res, byte(st[i].Val+'0'))
				st = append(st, st[i].Left)
				st = append(st, st[i].Right)
			}
		}
		res = append(res, '#')
		st = st[tempSize:]
	}
	return string(res)
}
func preSearch(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	preSearch(node.Left)
	preSearch(node.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numberOfBeams(bank []string) int {
	i := 0
	for i < len(bank) {
		zeroFlag := false
		for j := range bank[i] {
			if bank[i][j] == '1' {
				zeroFlag = true
				break
			}
		}
		if zeroFlag {
			i++
		} else {
			bank = append(bank[0:i], bank[i+1:]...)
		}
	}
	m := map[int]int{}
	fillMap(bank, m)
	res := 0
	for i := 0; i < len(bank)-1; i++ {
		res += (m[i] * m[i+1])
	}
	return res
}

func fillMap(bank []string, m map[int]int) {
	for i, str := range bank {
		for _, c := range str {
			if c == '1' {
				m[i]++
			}
		}
	}
}
