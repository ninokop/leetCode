package tree

func inorderTraversalByRecursion(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}

	nodes = append(nodes, inorderTraversalByRecursion(root.Left)...)
	nodes = append(nodes, root.Val)
	return append(nodes, inorderTraversalByRecursion(root.Right)...)
}

func inorderTraversalByStack(root *TreeNode) []int {
	s := newStack()
	out := make([]int, 0)

	p := root
	for p != nil || !s.Empty() {
		if p != nil {
			s.Push(p)
			p = p.Left
		} else {
			p = s.Top()
			s.Pop()

			out = append(out, p.Val)
			p = p.Right
		}
	}
	return out
}
