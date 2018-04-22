package tree

func minDepth(root *TreeNode) int {
	return minDepthOf(root, false)
}

func minDepthOf(root *TreeNode, hasBrother bool) int {
	if root == nil {
		if !hasBrother {
			return 0
		}
		return math.MaxInt64
	}

	return 1 + min(minDepthOf(root.Left, root.Right != nil),
		minDepthOf(root.Right, root.Left != nil))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left),
		maxDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
