package main

// 定义二叉树结构
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据层序遍历序列构建二叉树
func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// 构建左子节点
		if i < len(nums) {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
			i++
		}

		// 构建右子节点
		if i < len(nums) {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
			i++
		}
	}

	return root
}

// 辅助函数：层序遍历打印树结构（用于验证）
func levelOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func main() {
	// 题目给出的节点值序列
	//nums := []int{1, 2, 3, 5, null, 7, 8}

	// 构建二叉树
	//root := buildTree(nums)

	//fmt.Println(isCompleteTree(root))

	// 验证构建结果（层序遍历输出）
	//fmt.Println("层序遍历结果:", levelOrderTraversal(root))
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := []*TreeNode{root}
	for len(st) > 0 {
		tempSize := len(st)
		for i := 0; i < tempSize; i++ {
			if st[i] == nil {
				return false
			}
			if st[i].Right != nil && st[i].Left == nil {
			}
			if root.Left != nil {
				st = append(st, st[i].Left)
			}
			if root.Right != nil {
				st = append(st, st[i].Right)
			}
		}
		st = st[tempSize:]
	}
	return true
}
