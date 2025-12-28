package main

type Node struct {
	Val  int
	Next *Node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}
func dfs(root *TreeNode, targetVal int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == targetVal {
		return root
	}

	if root.Val > targetVal {
		//if root.Left==nil||root.Left.Val>targetVal{
		//	return nil
		//}
		return dfs(root.Left, targetVal)
	} else {
		//if root.Right==nil||root.Right.Val>targetVal{
		//	return nil
		//}
		return dfs(root.Right, targetVal)
	}
}

//func main() {
//
//	//1,2,3
//	//
//	//arr := make([]int, 1, 3)
//	//t := make([]int, len(arr))
//	//copy
//
//}

func reverseList(root *Node) *Node {
	var pre *Node
	cur := root
	for cur != nil {
		t := cur.Next
		cur.Next = pre
		pre = cur
		cur = t
	}
	return pre
}
