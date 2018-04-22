package tree

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBSTree(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTree(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	return root.Val > lower && root.Val < upper &&
		isValidBSTree(root.Left, lower, root.Val) &&
		isValidBSTree(root.Right, root.Val, upper)
}

func isValidBSTByInOrder(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	pre, cur := (*TreeNode)(nil), root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			if pre != nil && pre.Val >= cur.Val {
				return false
			}
			pre = cur
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return true
}
