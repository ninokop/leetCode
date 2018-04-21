package tree

func postorderTraversalByRecursion(root *TreeNode) []int {
	nodes := make([]int, 0)
	if root == nil {
		return nodes
	}
	nodes = append(nodes, postorderTraversalByRecursion(root.Left)...)
	nodes = append(nodes, postorderTraversalByRecursion(root.Right)...)
	return append(nodes, root.Val)
}

func postorderTraversalByStack(root *TreeNode) []int {
	s := newStack()
	out := make([]int, 0)

	p := root
	for {
		for p != nil {
			s.Push(p)
			p = p.Left
		}

		var q *TreeNode = nil
		for !s.Empty() {
			p = s.Top()
			s.Pop()

			if p.Right == q {
				out = append(out, p.Val)
				q = p
			} else {
				s.Push(p)
				p = p.Right
				break
			}
		}

		if s.Empty() {
			break
		}
	}
	return out
}
