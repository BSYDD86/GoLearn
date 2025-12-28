package main

//	   1
//	2    3
//
// 4  5  6  7
func main() {
	//   1
	// 2   3
	//4 5 6 7
	//moris

	//1242513637
	treeNode4 := &TreeNode{4, nil, nil}
	treeNode5 := &TreeNode{5, nil, nil}
	treeNode6 := &TreeNode{6, nil, nil}
	treeNode7 := &TreeNode{7, nil, nil}

	treeNode2 := &TreeNode{2, treeNode4, treeNode5}
	treeNode3 := &TreeNode{3, treeNode6, treeNode7}

	treeNode1 := &TreeNode{1, treeNode2, treeNode3}
	println(morrisLCA(treeNode1, treeNode4, treeNode2).Val)
}
func morrisLCA(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if preOrder(p.Left, p, q) != nil || preOrder(p.Right, p, q) != nil {
		return p
	}
	if preOrder(q.Left, p, q) != nil || preOrder(q.Right, p, q) != nil {
		return q
	}
	var (
		morrisRight *TreeNode
		lca         *TreeNode
	)
	left := preOrder(root, p, q)
	cur := root
	for cur != nil {
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
				if lca == nil {
					if checkRight(cur.Left, left) {
						if preOrder(left.Right, p, q) != nil {
							lca = left
						}
						left = cur
					}
				}
			}
		}
		cur = cur.Right
	}
	if lca != nil {
		return lca
	} else {
		return left
	}

}
func checkRight(root *TreeNode, target *TreeNode) bool {
	for root != nil {
		if root == target {
			return true
		}
		root = root.Right
	}
	return false
}

func preOrder(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var (
		morrisRight *TreeNode
		res         *TreeNode
	)
	cur := root

	for cur != nil {
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				if res == nil && (cur == p || cur == q) {
					res = cur
				}
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
			}
		} else {
			if res == nil && (cur == p || cur == q) {
				res = cur
			}
		}
		cur = cur.Right
	}
	return res
}

//func main() {
//
//	//   1
//	// 2   3
//	//4 5 6 7
//	//moris
//
//	//1242513637
//	treeNode4 := &TreeNode{4, nil, nil}
//	treeNode5 := &TreeNode{5, nil, nil}
//	treeNode6 := &TreeNode{6, nil, nil}
//	treeNode7 := &TreeNode{7, nil, nil}
//
//	treeNode2 := &TreeNode{2, treeNode4, treeNode5}
//	treeNode3 := &TreeNode{3, treeNode6, treeNode7}
//
//	treeNode1 := &TreeNode{1, treeNode2, treeNode3}
//
//	morrisTraversal(treeNode1)
//
//}
//func morrisTraversal(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	var (
//		morrisRight *TreeNode
//	)
//	cur := root
//	for cur != nil {
//		fmt.Println(cur.Val)
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
//			}
//		}
//		cur = cur.Right
//	}
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
