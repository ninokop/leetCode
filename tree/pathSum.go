package tree

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	return hasPathSum(root.Left, sum-root.Val) ||
		hasPathSum(root.Right, sum-root.Val)
}

func pathSum(root *TreeNode, sum int) [][]int {
	out := make([][]int, 0)
	path := make([]int, 0)
	pathSumOf(root, sum, path, &out)
	return out
}

func pathSumOf(root *TreeNode, sum int, path []int, out *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	fmt.Println(path, sum)
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			*out = append(*out, copy(path))

		}
	}

	pathSumOf(root.Left, sum-root.Val, path, out)
	pathSumOf(root.Right, sum-root.Val, path, out)

	path = path[:len(path)-1]
}

func copy(path []int) []int {
	out := make([]int, len(path))
	for i := range path {
		out[i] = path[i]
	}
	return out
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	sum := root.Val
	if left > 0 {
		sum += left
	}

	if right > 0 {
		sum += right
	}

	*maxSum = max(*maxSum, sum)

	if child := max(right, left); child > 0 {
		return child + root.Val
	} else {
		return root.Val
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sumNumbers(root *TreeNode) int {
	return dfsPathSum(root, 0)
}

func dfsPathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return sum*10 + root.Val
	}

	return dfsPathSum(root.Left, sum*10+root.Val) +
		dfsPathSum(root.Right, sum*10+root.Val)
}
