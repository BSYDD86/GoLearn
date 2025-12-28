package main

func main() {

}
func generateByArr(arr []int, idx int) *TreeNode {
	if idx >= len(arr) {
		return nil
	}
	if arr[idx] == -1 {
		return nil
	}
	root := &TreeNode{Val: arr[idx], Left: generateByArr(arr, 2*idx+1), Right: generateByArr(arr, 2*idx+2)}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
