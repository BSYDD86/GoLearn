package main

// import "fmt"
//
//	func main() {
//		//   1
//		// 2   3
//		//4 5 6 7
//		//moris
//
//		//1242513637
//
//		//preOrder 1245 367
//		//inOrder 4251 637
//		//postOrder 452 673 1
//		treeNode4 := &TreeNode{4, nil, nil}
//		treeNode5 := &TreeNode{5, nil, nil}
//		treeNode6 := &TreeNode{6, nil, nil}
//		treeNode7 := &TreeNode{7, nil, nil}
//
//		treeNode2 := &TreeNode{2, treeNode4, treeNode5}
//		treeNode3 := &TreeNode{3, treeNode6, treeNode7}
//
//		treeNode1 := &TreeNode{1, treeNode2, treeNode3}
//
//		res = []int{}
//		morris(treeNode1)
//		collectLeftTreeRightChildren(treeNode1)
//		fmt.Println(res)
//	}
//func morris(root *TreeNode) {
//	if root == nil {
//		return
//	}
//
//	cur := root
//	var morrisRight *TreeNode
//	for cur != nil {
//
//		morrisRight = cur.Left
//		if morrisRight != nil {
//			for morrisRight.Right != nil && morrisRight.Right != cur {
//				morrisRight = morrisRight.Right
//			}
//			if morrisRight.Right == nil {
//				morrisRight.Right = cur
//				cur = cur.Left
//				continue
//			} else {
//				morrisRight.Right = nil
//				collectLeftTreeRightChildren(cur.Left)
//			}
//		}
//		cur = cur.Right
//	}
//}

//func collectLeftTreeRightChildren(root *TreeNode) {
//	temp := reverseTree(root)
//	tempCollect := temp
//	for tempCollect != nil {
//		res = append(res, tempCollect.Val)
//		tempCollect = tempCollect.Right
//	}
//	reverseTree(temp)
//}
//
//func reverseTree(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	cur := root
//	var pre *TreeNode
//	for cur != nil {
//		temp := cur.Right
//		cur.Right = pre
//		pre = cur
//		cur = temp
//	}
//	return pre
//}

//
//var (
//	res []int
//)
